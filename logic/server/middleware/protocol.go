package middleware

import (
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request, Extractor)
type MiddleOpFunc func(HandlerFunc) HandlerFunc

type GetHandler = func(http.ResponseWriter, *http.Request, Extractor) Handler
