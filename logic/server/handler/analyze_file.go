package handler

import (
	"code-comment-analyzer/ccanalyzer_client"
	"code-comment-analyzer/data"
	"code-comment-analyzer/server/middleware"
	"net/http"
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

}
