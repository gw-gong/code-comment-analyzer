package public

import (
	"net/http"

	"code-comment-analyzer/config"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
	"code-comment-analyzer/util"
)

type GetReadme struct {
	w http.ResponseWriter
	r *http.Request
}

func NewGetReadme() middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &GetReadme{
			w: w,
			r: r,
		}
	}
}

func (gr *GetReadme) Handle() {
	fileContent, err := util.ReadFileContentByPath(config.Cfg.ReadmePath)
	if err != nil {
		protocol.HttpResponseFail(gr.w, http.StatusInternalServerError, protocol.ErrorCodeInternalServerError, "获取说明文档失败")
		return
	}

	protocol.HttpResponseSuccess(gr.w, http.StatusOK, "Success", protocol.WithData(fileContent))
}
