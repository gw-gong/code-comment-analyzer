package middleware

import (
	"log"
	"net/http"
)

func RegisterRouter(method string, pattern string, getHandler GetHandler, middleOps ...MiddleOpFunc) {
	if mux == nil {
		panic("mux is nil")
	}
	handlerFunc := defaultMiddleOp(getHandler)
	for i := len(middleOps) - 1; i >= 0; i-- {
		handlerFunc = middleOps[i](handlerFunc)
	}
	formatHandlerFunc := func(w http.ResponseWriter, r *http.Request) {
		switch method {
		case Get:
			handlerFunc = enforceGet(handlerFunc)
		case Post:
			handlerFunc = enforcePost(handlerFunc)
		default:
			log.Printf("invalid method: %s", method)
		}
		handlerFunc(w, r, nil)
	}
	mux.HandleFunc(pattern, formatHandlerFunc)
}

func defaultMiddleOp(getHandler GetHandler) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, extractor Extractor) {
		h := getHandler(w, r, extractor)
		h.Handle()
	}
}
