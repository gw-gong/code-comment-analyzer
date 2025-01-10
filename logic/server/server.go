package server

import (
	"fmt"
	"log"
	"net/http"

	"code-comment-analyzer/ccanalyzer_client"
	"code-comment-analyzer/data"
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

func (s *Server) RegisterRouters(register *data.DataManagerRegistry, ccanalyzer ccanalyzer_client.CCAnalyzer) {
	middleware.RegisterSessionManager(register.GetSessionManager())
	middleware.RegisterRouter(s.mux, "/test", handler.NewTestXXX(register, ccanalyzer), middleware.EnforceGet, middleware.AuthenticateUser)
	middleware.RegisterRouter(s.mux, "/test_login", handler.NewLogin(register), middleware.EnforcePost)
}

func (s *Server) Listen(host, port string) {
	log.Println("Listening on " + host + ":" + port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), s.mux)
	if err != nil {
		panic(err)
	}
}
