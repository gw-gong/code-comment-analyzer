package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"code-comment-analyzer/ccanalyzer_client"
	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
)

type AnalyzeFile struct {
	w          http.ResponseWriter
	r          *http.Request
	extractor  middleware.Extractor
	registry   *data.DataManagerRegistry
	ccanalyzer ccanalyzer_client.CCAnalyzer
}

func NewAnalyzeFile(registry *data.DataManagerRegistry, ccanalyzer ccanalyzer_client.CCAnalyzer) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &AnalyzeFile{
			w:          w,
			r:          r,
			extractor:  extractor,
			registry:   registry,
			ccanalyzer: ccanalyzer,
		}
	}
}

func (af *AnalyzeFile) Handle() {
	requestData, err := af.DecodeRequest()
	if err != nil {
		return
	}

	if protocol.IsLanguageSupported(requestData.Language) == false {
		protocol.HttpResponseFail(af.w, http.StatusBadRequest, protocol.ErrorCodeLanguageNotSupported, "Language not supported")
		return
	}

	analyzedData, err := af.ccanalyzer.AnalyzeFileContent(requestData.Language, requestData.FileContent)
	if err != nil {
		protocol.HttpResponseFail(af.w, http.StatusInternalServerError, protocol.ErrorCodeRPCCallFail, fmt.Sprintf("%v", err))
		return
	}
	protocol.HttpResponseSuccess(af.w, http.StatusOK, "Success", analyzedData, requestData.Language)
}

func (af *AnalyzeFile) DecodeRequest() (*protocol.AnalyzeFileRequest, error) {
	var requestData protocol.AnalyzeFileRequest
	err := json.NewDecoder(af.r.Body).Decode(&requestData)
	if err != nil {
		log.Println("Failed to parse JSON body:", err)
		protocol.HttpResponseFail(af.w, http.StatusBadRequest, protocol.ErrorCodeParseRequestFailed, "Invalid JSON format")
		return nil, err
	}
	return &requestData, nil
}
