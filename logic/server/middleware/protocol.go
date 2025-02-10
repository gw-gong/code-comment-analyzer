package middleware

import (
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request, extractor Extractor)
type MiddleOpFunc func(rg *routerGroup, handlerFunc HandlerFunc) HandlerFunc

type GetHandler func(w http.ResponseWriter, r *http.Request, extractor Extractor) Handler

type Handler interface {
	Handle()
}

const (
	Get  = "GET"
	Post = "POST"
)
