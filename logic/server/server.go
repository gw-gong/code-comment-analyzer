package server

import (
	"code-comment-analyzer/data"
	"code-comment-analyzer/server/handler"
	"code-comment-analyzer/server/middleware"
	"fmt"
	"log"
	"net/http"
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

func (s *Server) RegisterRouters(register *data.DataManagerRegistry) {
	middleware.RegisterRouter(s.mux, "/test", handler.NewTestXXX(register), middleware.Get, middleware.AuthenticateForUser)
	middleware.RegisterRouter(s.mux, "/test_login", handler.NewLogin(register), middleware.Post)
}

func (s *Server) Listen(host, port string) {
	log.Println("Listening on " + host + ":" + port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), s.mux)
	if err != nil {
		panic(err)
	}
}
