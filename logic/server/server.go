package server

import (
	"fmt"
	"log"
	"net/http"

	"code-comment-analyzer/server/handler"
	"code-comment-analyzer/server/middleware"
)

type Server struct {
	mux *http.ServeMux
}

func NewHTTPServer() *Server {
	s := &Server{
		mux: http.NewServeMux(),
	}
	return s
}

func (s *Server) RegisterRouters() {
	middleware.RegisterRouter(s.mux, "/test", handler.Test, middleware.AuthenticationForUser)
	middleware.RegisterRouter(s.mux, "/test_login", handler.TestLogin)
}

func (s *Server) Listen(host, port string) {
	log.Println("Listening on " + host + ":" + port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), s.mux)
	if err != nil {
		panic(err)
	}
}
