package middleware

import (
	"net/http"
)

func RegisterRouter(mux *http.ServeMux, pattern string, getHandler GetHandler, middleOps ...MiddleOpFunc) {
	if mux == nil {
		panic("mux is nil")
	}
	handlerFunc := defaultMiddleOp(getHandler)
	for i := len(middleOps) - 1; i >= 0; i-- {
		handlerFunc = middleOps[i](handlerFunc)
	}
	formatHandlerFunc := func(w http.ResponseWriter, r *http.Request) {
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
