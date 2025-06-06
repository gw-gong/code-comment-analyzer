package public

import (
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

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

	rootNode := util.BuildDirectoryTree(destDir, destDir, projectStorageName)

	protocol.HttpResponseSuccess(u.w, http.StatusOK, "文件已解压", protocol.WithData(rootNode.Children[0]))

	go util.WithRecover(func() { u.recordProjectUpload(destDir) })
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
