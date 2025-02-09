package public

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"code-comment-analyzer/config"
	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
	"code-comment-analyzer/server/server_info"
	"code-comment-analyzer/util"
)

type ReadFile struct {
	w         http.ResponseWriter
	r         *http.Request
	extractor middleware.Extractor
	registry  *data.DataManagerRegistry
}

func NewReadFile(registry *data.DataManagerRegistry) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &ReadFile{
			w:         w,
			r:         r,
			extractor: extractor,
			registry:  registry,
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

	curProjectRunningPath := server_info.GetServerRunningPath()
	absolutePath := filepath.Join(curProjectRunningPath, config.Cfg.FileStoragePath.Projects, requestData.Path)

	if _, err = os.Stat(absolutePath); os.IsNotExist(err) {
		protocol.HttpResponseFail(rf.w, http.StatusNotFound, protocol.ErrorCodeFileNotFound, "文件不存在")
		return
	} else if err != nil {
		protocol.HttpResponseFail(rf.w, http.StatusInternalServerError, protocol.ErrorCodeInternalServerError, "检查文件状态失败")
		return
	}

	file, err := os.Open(absolutePath)
	if err != nil {
		protocol.HttpResponseFail(rf.w, http.StatusInternalServerError, protocol.ErrorCodeInternalServerError, "打开文件失败")
		return
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		protocol.HttpResponseFail(rf.w, http.StatusInternalServerError, protocol.ErrorCodeInternalServerError, "读取文件失败")
		return
	}

	fileSuffix := filepath.Ext(absolutePath)
	language := util.FileSuffixToLanguage(fileSuffix)

	response := protocol.ReadFileResponse{
		FileContent: string(fileContent),
	}

	protocol.HttpResponseSuccess(rf.w, http.StatusOK, "Success", response, language)
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
