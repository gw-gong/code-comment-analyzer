package handler

import (
	"net/http"

	"code-comment-analyzer/data"
	"code-comment-analyzer/server/jwt"
	"code-comment-analyzer/server/middleware"
)

type Login struct {
	w         http.ResponseWriter
	r         *http.Request
	extractor middleware.Extractor
	registry  *data.DataManagerRegistry
}

func NewLogin(registry *data.DataManagerRegistry) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &Login{
			w:         w,
			r:         r,
			extractor: extractor,
			registry:  registry,
		}
	}
}

func (l *Login) Handle() {
	// ....
	jwt.AuthorizeUserToken(123456, l.w, l.registry.GetSessionManager())
}
