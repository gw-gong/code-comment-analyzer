package public

import (
	"io"
	"log"
	"net/http"
	"path/filepath"

	"code-comment-analyzer/config"
	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
	"code-comment-analyzer/util"
)

type File2String struct {
	w         http.ResponseWriter
	r         *http.Request
	extractor middleware.Extractor
	registry  *data.DataManagerRegistry
}

func NewFile2String(registry *data.DataManagerRegistry) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &File2String{
			w:         w,
			r:         r,
			extractor: extractor,
			registry:  registry,
		}
	}
}

func (f2s *File2String) Handle() {
	fileBytes, language, err := f2s.decodeRequest()
	if err != nil {
		return
	}

	fileContent := string(fileBytes)
	response := &protocol.File2StringResponse{
		Language:    language,
		FileContent: fileContent,
	}
	protocol.HttpResponseSuccess(f2s.w, http.StatusOK, "已读取", protocol.WithData(response))

	go util.WithRecover(func() { f2s.recordFileUpload(language, fileContent) })
}

func (f2s *File2String) decodeRequest() (fileContent []byte, fileType string, err error) {
	maxFileSize := config.Cfg.MaxFileSize
	err = f2s.r.ParseMultipartForm(maxFileSize << 20)
	if err != nil {
		protocol.HttpResponseFail(f2s.w, http.StatusBadRequest, protocol.ErrorCodeFileTooLarge, "file too large")
		return nil, "", err
	}

	file, header, err := f2s.r.FormFile(protocol.MultipartFormKeyFile)
	if err != nil {
		protocol.HttpResponseFail(f2s.w, http.StatusBadRequest, protocol.ErrorCodeFileNotFound, "file not found")
		return nil, "", err
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		protocol.HttpResponseFail(f2s.w, http.StatusBadRequest, protocol.ErrorCodeInternalServerError, "Unable to read file")
		return nil, "", err
	}

	fileName := header.Filename
	fileSuffix := filepath.Ext(fileName)
	language := util.FileSuffixToLanguage(fileSuffix)
	return fileBytes, language, nil
}

func (f2s *File2String) recordFileUpload(language, fileContent string) {
	if isUserLoggedIn, err := f2s.extractor.IsUserLoggedIn(); err != nil || !isUserLoggedIn {
		return
	}
	userID, err := f2s.extractor.GetUserId()
	if err != nil {
		return
	}
	log.Println("recordFileUpload|userID", userID)

	om := f2s.registry.GetOperationManager()
	err = om.RecordFileUpload(userID, language, fileContent)
	if err != nil {
		log.Println("recordFileUpload|RecordFileUpload|err", err)
	}
}
