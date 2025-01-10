package middleware

import (
	"code-comment-analyzer/protocol"
	"fmt"
	"net/http"
)

const (
	MethodGet  = "GET"
	MethodPost = "POST"
)

func Get(handlerFunc HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, extractor Extractor) {
		if r.Method != MethodGet {
			protocol.HandleError(w, protocol.ErrorCodeMustBeGet, fmt.Errorf("request method must be %s", MethodGet))
			return
		}
		handlerFunc(w, r, extractor)
	}
}

func Post(handlerFunc HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, extractor Extractor) {
		if r.Method != MethodPost {
			protocol.HandleError(w, protocol.ErrorCodeMustBeGet, fmt.Errorf("request method must be %s", MethodPost))
			return
		}
		handlerFunc(w, r, extractor)
	}
}
