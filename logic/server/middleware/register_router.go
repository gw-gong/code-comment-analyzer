package middleware

import "net/http"

func RegisterRouter(mux *http.ServeMux, pattern string, handlerFunc HandlerFunc, middleOps ...MiddleOpFunc) {
	if mux == nil {
		panic("mux is nil")
	}
	finalHandler := handlerFunc
	for _, middleOp := range middleOps {
		finalHandler = middleOp(finalHandler)
	}
	mux.HandleFunc(pattern, finalHandler)
}
