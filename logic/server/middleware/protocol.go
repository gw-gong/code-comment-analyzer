package middleware

import "net/http"

type HandlerFunc func(http.ResponseWriter, *http.Request)
type MiddleOpFunc func(HandlerFunc) HandlerFunc
