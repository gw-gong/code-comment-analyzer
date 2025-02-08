package public

import (
	"code-comment-analyzer/config"
	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
	"code-comment-analyzer/util"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type UploadAndGetTree struct {
	w         http.ResponseWriter
	r         *http.Request
	extractor middleware.Extractor
	registry  *data.DataManagerRegistry
}

func NewUploadAndGetTree(registry *data.DataManagerRegistry) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &File2String{
			w:         w,
			r:         r,
			extractor: extractor,
			registry:  registry,
		}
	}
}

func (uagt *UploadAndGetTree) Handle() {
	file, header, err := uagt.decodeRequest()
	if err != nil {
		return
	}
	defer file.Close()

	projectsRootPath := config.Cfg.FileStoragePath.Projects
	projectName := strings.TrimSuffix(header.Filename, protocol.FileSuffixZIP)
	destDir := filepath.Join(projectsRootPath, projectName)
	if err = os.MkdirAll(destDir, 0755); err != nil {
		protocol.HttpResponseFail(uagt.w, http.StatusInternalServerError, protocol.ErrorCodeCreatePathFailed, "创建目录失败")
		return
	}

	tempZipPath := filepath.Join(destDir, header.Filename)
	defer os.Remove(tempZipPath)

	out, err := os.Create(tempZipPath)
	if err != nil {
		protocol.HttpResponseFail(uagt.w, http.StatusInternalServerError, protocol.ErrorCodeSaveFileFailed, "创建临时文件失败")
		return
	}
	defer out.Close()

	if _, err = io.Copy(out, file); err != nil {
		protocol.HttpResponseFail(uagt.w, http.StatusInternalServerError, protocol.ErrorCodeSaveFileFailed, "保存文件失败")
		return
	}

	if err = util.Unzip(tempZipPath, destDir); err != nil {
		protocol.HttpResponseFail(uagt.w, http.StatusInternalServerError, protocol.ErrorCodeUnzipFailed, "解压文件失败")
		return
	}

	rootNode := uagt.buildDirectoryTree(destDir, destDir)
	response := protocol.FileNode{
		Label:    projectName,
		Children: rootNode.Children,
	}

	protocol.HttpResponseSuccess(uagt.w, http.StatusOK, "文件已解压", response)
}

func (uagt *UploadAndGetTree) decodeRequest() (file multipart.File, header *multipart.FileHeader, err error) {
	maxProjectSize := config.Cfg.MaxProjectSize
	err = uagt.r.ParseMultipartForm(maxProjectSize << 20)
	if err != nil {
		protocol.HttpResponseFail(uagt.w, http.StatusBadRequest, protocol.ErrorCodeFileTooLarge, "file too large")
		return nil, nil, err
	}

	file, header, err = uagt.r.FormFile(protocol.MultipartFormKeyFile)
	if err != nil {
		protocol.HttpResponseFail(uagt.w, http.StatusBadRequest, protocol.ErrorCodeFileNotFound, "file not found")
		return nil, nil, err
	}

	// todo 判断是不是.zip，不是zip，直接file.close()

	return
}

func (uagt *UploadAndGetTree) buildDirectoryTree(rootPath, currentPath string) protocol.FileNode {
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
			child := uagt.buildDirectoryTree(rootPath, fullPath)
			node.Children = append(node.Children, child)
		} else {
			node.Children = append(node.Children, protocol.FileNode{
				Label: entry.Name(),
				Value: filepath.Join("file_storage", "projects", relPath),
			})
		}
	}

	return node
}
