package public

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"code-comment-analyzer/config"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
	"code-comment-analyzer/util"
)

type ReadFile struct {
	w http.ResponseWriter
	r *http.Request
}

func NewReadFile() middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &ReadFile{
			w: w,
			r: r,
		}
	}
}

func (rf *ReadFile) Handle() {
	requestData, err := rf.decodeRequest()
	if err != nil {
		return
	}

	if requestData.Path == "" {
		protocol.HttpResponseFail(rf.w, http.StatusBadRequest, protocol.ErrorCodeParamError, "文件路径为空")
		return
	}

	realPath := filepath.Join(config.Cfg.FileStoragePath.Projects, requestData.Path)

	if _, err = os.Stat(realPath); os.IsNotExist(err) {
		protocol.HttpResponseFail(rf.w, http.StatusNotFound, protocol.ErrorCodeFileNotFound, "文件不存在")
		return
	} else if err != nil {
		protocol.HttpResponseFail(rf.w, http.StatusInternalServerError, protocol.ErrorCodeInternalServerError, "检查文件状态失败")
		return
	}

	fileContent, err := util.ReadFileContentByPath(realPath)
	if err != nil {
		protocol.HttpResponseFail(rf.w, http.StatusInternalServerError, protocol.ErrorCodeInternalServerError, fmt.Sprintf("%v", err))
		return
	}

	fileSuffix := filepath.Ext(realPath)
	language := util.FileSuffixToLanguage(fileSuffix)

	response := protocol.ReadFileResponse{
		FileContent: fileContent,
	}

	protocol.HttpResponseSuccess(rf.w, http.StatusOK, "Success", protocol.WithData(response), protocol.WithLanguage(language))
}

func (rf *ReadFile) decodeRequest() (*protocol.ReadFileRequest, error) {
	var requestData protocol.ReadFileRequest
	err := json.NewDecoder(rf.r.Body).Decode(&requestData)
	if err != nil {
		log.Println("Failed to parse JSON body:", err)
		protocol.HttpResponseFail(rf.w, http.StatusBadRequest, protocol.ErrorCodeParseRequestFailed, "Invalid JSON format")
		return nil, err
	}
	return &requestData, nil
}
