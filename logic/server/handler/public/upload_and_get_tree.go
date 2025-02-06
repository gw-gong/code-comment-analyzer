package public

import (
	"code-comment-analyzer/config"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/util"
	"io"
	"net/http"
	"path/filepath"

	"code-comment-analyzer/data"
	"code-comment-analyzer/server/middleware"
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

}

func (uagt *UploadAndGetTree) decodeRequest() (fileContent []byte, fileType string, err error) {
	maxFileSize := config.Cfg.MaxFileSize
	err = uagt.r.ParseMultipartForm(maxFileSize << 20) // 限制上传文件大小为 10MB
	if err != nil {
		protocol.HttpResponseFail(uagt.w, http.StatusBadRequest, protocol.ErrorCodeFileTooLarge, "file too large")
		return nil, "", err
	}

	file, header, err := uagt.r.FormFile("file") // "file" 是表单中的 key
	if err != nil {
		protocol.HttpResponseFail(uagt.w, http.StatusBadRequest, protocol.ErrorCodeFileNotFound, "file not found")
		return nil, "", err
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		protocol.HttpResponseFail(uagt.w, http.StatusBadRequest, protocol.ErrorCodeInternalServerError, "Unable to read file")
		return nil, "", err
	}

	fileName := header.Filename
	fileSuffix := filepath.Ext(fileName)
	language := util.FileSuffixToLanguage(fileSuffix)
	return fileBytes, language, nil
}
