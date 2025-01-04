package server

import (
	"fmt"
	"log"
	"net/http"

	"code-comment-analyzer/server/handler"
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
	s.mux.HandleFunc("/test", handler.Test)
}

func (s *Server) Listen(host, port string) {
	log.Println("Listening on " + host + ":" + port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), s.mux)
	if err != nil {
		panic(err)
	}
}
