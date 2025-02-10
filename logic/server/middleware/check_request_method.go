package middleware

import (
	"fmt"
	"net/http"

	"code-comment-analyzer/protocol"
)

func enforceGet(handlerFunc HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, extractor Extractor) {
		if r.Method != Get {
			protocol.HttpResponseFail(w, http.StatusInternalServerError, protocol.ErrorCodeMustBeGet, fmt.Sprintf("request method must be %s", Get))
			return
		}
		handlerFunc(w, r, extractor)
	}
}

func enforcePost(handlerFunc HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, extractor Extractor) {
		if r.Method != Post {
			protocol.HttpResponseFail(w, http.StatusInternalServerError, protocol.ErrorCodeMustBePost, fmt.Sprintf("request method must be %s", Post))
			return
		}
		handlerFunc(w, r, extractor)
	}
}
