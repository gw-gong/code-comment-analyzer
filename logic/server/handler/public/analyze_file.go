package public

import (
	"encoding/json"
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
	requestData, err := af.decodeRequest()
	if err != nil {
		return
	}

	if protocol.IsLanguageSupported(requestData.Language) == false {
		protocol.HttpResponseFail(af.w, http.StatusBadRequest, protocol.ErrorCodeLanguageNotSupported, "Language not supported")
		return
	}

	var analyzedData protocol.AnalyzeFileResponse
	analyzedData, err = af.ccanalyzer.AnalyzeFileContent(requestData.Language, requestData.FileContent)
	if err != nil {
		log.Println("Failed to analyze file content:", err)
		protocol.HttpResponseFail(af.w, http.StatusInternalServerError, protocol.ErrorCodeAnalyzeFileFailed, "Failed to analyze file content")
		return
	}

	protocol.HttpResponseSuccess(af.w, http.StatusOK, "Success", protocol.WithData(analyzedData), protocol.WithLanguage(requestData.Language))
}

func (af *AnalyzeFile) decodeRequest() (*protocol.AnalyzeFileRequest, error) {
	var requestData protocol.AnalyzeFileRequest
	err := json.NewDecoder(af.r.Body).Decode(&requestData)
	if err != nil {
		log.Println("Failed to parse JSON body:", err)
		protocol.HttpResponseFail(af.w, http.StatusBadRequest, protocol.ErrorCodeParseRequestFailed, "Invalid JSON format")
		return nil, err
	}
	return &requestData, nil
}
