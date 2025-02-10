package server

import (
	"fmt"
	"log"
	"net/http"

	"code-comment-analyzer/ccanalyzer_client"
	"code-comment-analyzer/data"
	"code-comment-analyzer/server/handler/public"
	"code-comment-analyzer/server/handler/user"
	m "code-comment-analyzer/server/middleware"
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

func (s *Server) RegisterRouters(registry *data.DataManagerRegistry, ccanalyzer ccanalyzer_client.CCAnalyzer) {
	m.RegisterMux(s.mux)
	m.RegisterSessionManager(registry.GetSessionManager())

	m.RegisterRouter(m.Post, "/public/upload_file2string/", public.NewFile2String(registry), m.CheckLoginStatus)
	m.RegisterRouter(m.Post, "/public/analyze_file/", public.NewAnalyzeFile(registry, ccanalyzer))
	m.RegisterRouter(m.Post, "/public/upload_and_get_tree/", public.NewUploadAndGetTree(registry), m.CheckLoginStatus)
	m.RegisterRouter(m.Post, "/public/read_file/", public.NewReadFile())
	m.RegisterRouter(m.Get, "/public/get_readme/", public.NewGetReadme())

	m.RegisterRouter(m.Post, "/user/login/", user.NewLogin(registry))
	m.RegisterRouter(m.Get, "/user/logout/", user.NewLogout(registry), m.CheckLoginStatus)
	m.RegisterRouter(m.Post, "/user/sign_up/", user.NewSignup(registry))
	m.RegisterRouter(m.Get, "/user/get_user_info/", user.NewGetUserInfo(registry), m.AuthenticateUser)
	m.RegisterRouter(m.Get, "/user/get_user_profile_picture/", user.NewGetUserProfilePicture(registry), m.CheckLoginStatus)
}

func (s *Server) Listen(host, port string) {
	log.Println("Listening on " + host + ":" + port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), s.mux)
	if err != nil {
		panic(err)
	}
}
