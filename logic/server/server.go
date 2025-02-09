package server

import (
	"fmt"
	"log"
	"net/http"

	"code-comment-analyzer/ccanalyzer_client"
	"code-comment-analyzer/data"
	"code-comment-analyzer/server/handler"
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
	m.RegisterSessionManager(registry.GetSessionManager())
	m.RegisterRouter(s.mux, "/test/", handler.NewTestXXX(registry, ccanalyzer), m.EnforceGet, m.AuthenticateUser)

	m.RegisterRouter(s.mux, "/public/upload_file2string/", public.NewFile2String(registry), m.EnforcePost, m.CheckLoginStatus)
	m.RegisterRouter(s.mux, "/public/analyze_file/", public.NewAnalyzeFile(registry, ccanalyzer), m.EnforcePost)
	m.RegisterRouter(s.mux, "/public/upload_and_get_tree/", public.NewUploadAndGetTree(registry), m.EnforcePost, m.CheckLoginStatus)
	m.RegisterRouter(s.mux, "/public/read_file/", public.NewReadFile(), m.EnforcePost)
	m.RegisterRouter(s.mux, "/public/get_readme/", public.NewGetReadme(), m.EnforceGet)

	m.RegisterRouter(s.mux, "/user/login/", user.NewLogin(registry), m.EnforcePost)
	m.RegisterRouter(s.mux, "/user/logout/", user.NewLogout(registry), m.EnforceGet, m.CheckLoginStatus)
	m.RegisterRouter(s.mux, "/user/sign_up/", user.NewSignup(registry), m.EnforcePost)
	m.RegisterRouter(s.mux, "/user/get_user_info/", user.NewGetUserInfo(registry), m.EnforceGet, m.AuthenticateUser)
	m.RegisterRouter(s.mux, "/user/get_user_profile_picture/", user.NewGetUserProfilePicture(registry), m.EnforceGet, m.CheckLoginStatus)
}

func (s *Server) Listen(host, port string) {
	log.Println("Listening on " + host + ":" + port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), s.mux)
	if err != nil {
		panic(err)
	}
}
