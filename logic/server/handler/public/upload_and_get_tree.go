package public

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"code-comment-analyzer/config"
	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
	"code-comment-analyzer/util"
)

type UploadAndGetTree struct {
	w         http.ResponseWriter
	r         *http.Request
	extractor middleware.Extractor
	registry  *data.DataManagerRegistry
}

func NewUploadAndGetTree(registry *data.DataManagerRegistry) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &UploadAndGetTree{
			w:         w,
			r:         r,
			extractor: extractor,
			registry:  registry,
		}
	}
}

func (u *UploadAndGetTree) Handle() {
	file, header, err := u.decodeRequest()
	if err != nil {
		return
	}
	defer file.Close()

	projectName := strings.TrimSuffix(header.Filename, protocol.FileSuffixZIP)
	projectsStorageRootPath := config.Cfg.FileStoragePath.Projects
	projectStorageName := util.GenerateUUIDProjectName()
	destDir := filepath.Join(projectsStorageRootPath, projectStorageName)
	if err = os.MkdirAll(destDir, 0755); err != nil {
		protocol.HttpResponseFail(u.w, http.StatusInternalServerError, protocol.ErrorCodeCreatePathFailed, "创建目录失败")
		return
	}

	tempZipPath := filepath.Join(destDir, header.Filename)

	out, err := os.Create(tempZipPath)
	if err != nil {
		protocol.HttpResponseFail(u.w, http.StatusInternalServerError, protocol.ErrorCodeSaveFileFailed, "创建临时文件失败")
		return
	}
	defer out.Close()

	if _, err = io.Copy(out, file); err != nil {
		protocol.HttpResponseFail(u.w, http.StatusInternalServerError, protocol.ErrorCodeSaveFileFailed, "保存文件失败")
		return
	}

	if err = util.Unzip(tempZipPath, destDir); err != nil {
		protocol.HttpResponseFail(u.w, http.StatusInternalServerError, protocol.ErrorCodeUnzipFailed, "解压文件失败")
		return
	}
	os.Remove(tempZipPath)

	rootNode := u.buildDirectoryTree(destDir, destDir, projectStorageName)
	response := protocol.FileNode{
		Label:    projectName,
		Children: rootNode.Children,
	}

	protocol.HttpResponseSuccess(u.w, http.StatusOK, "文件已解压", response)

	go u.recordProjectUpload(destDir)
}

func (u *UploadAndGetTree) decodeRequest() (file multipart.File, header *multipart.FileHeader, err error) {
	maxProjectSize := config.Cfg.MaxProjectSize
	err = u.r.ParseMultipartForm(maxProjectSize << 20)
	if err != nil {
		protocol.HttpResponseFail(u.w, http.StatusBadRequest, protocol.ErrorCodeFileTooLarge, "file too large")
		return nil, nil, err
	}

	file, header, err = u.r.FormFile(protocol.MultipartFormKeyFile)
	if err != nil {
		protocol.HttpResponseFail(u.w, http.StatusBadRequest, protocol.ErrorCodeFileNotFound, "file not found")
		return nil, nil, err
	}

	// todo 判断是不是.zip，不是zip，直接file.close()

	return
}

func (u *UploadAndGetTree) buildDirectoryTree(currentPath, rootPath, projectStorageName string) protocol.FileNode {
	node := protocol.FileNode{
		Label: filepath.Base(currentPath),
	}

	entries, err := os.ReadDir(currentPath)
	if err != nil {
		return node
	}

	for _, entry := range entries {
		fullPath := filepath.Join(currentPath, entry.Name())
		relPath, _ := filepath.Rel(rootPath, fullPath)
		relPath = filepath.ToSlash(relPath) // 统一使用斜杠

		if entry.IsDir() {
			child := u.buildDirectoryTree(fullPath, rootPath, projectStorageName)
			node.Children = append(node.Children, child)
		} else {
			node.Children = append(node.Children, protocol.FileNode{
				Label: entry.Name(),
				Value: filepath.Join(projectStorageName, relPath),
			})
		}
	}

	return node
}

func (u *UploadAndGetTree) recordProjectUpload(projectUrl string) {
	if isUserLoggedIn, err := u.extractor.IsUserLoggedIn(); err != nil || !isUserLoggedIn {
		return
	}
	userID, err := u.extractor.GetUserId()
	if err != nil {
		return
	}
	log.Println("recordProjectUpload|userID", userID)

	om := u.registry.GetOperationManager()
	err = om.RecordProjectUpload(userID, projectUrl)
	if err != nil {
		log.Println("recordProjectUpload|RecordProjectUpload|err:", err)
	}
}
