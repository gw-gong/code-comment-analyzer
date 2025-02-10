package middleware

import (
	"code-comment-analyzer/data/redis"
	"net/http"
)

var (
	mux            *http.ServeMux
	sessionManager redis.SessionManager
)

func RegisterMux(m *http.ServeMux) {
	mux = m
}

func RegisterSessionManager(s redis.SessionManager) {
	sessionManager = s
}
