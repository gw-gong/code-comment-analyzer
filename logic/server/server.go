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
	sm := registry.GetSessionManager()

	publicGroup := m.NewRouterGroup("/public/", s.mux, m.WithSessionManager(sm))
	publicGroup.Post("upload_file2string/", public.NewFile2String(registry), m.CheckLoginStatus)
	publicGroup.Post("analyze_file/", public.NewAnalyzeFile(registry, ccanalyzer))
	publicGroup.Post("upload_and_get_tree/", public.NewUploadAndGetTree(registry), m.CheckLoginStatus)
	publicGroup.Post("read_file/", public.NewReadFile())
	publicGroup.Get("get_readme/", public.NewGetReadme())

	userGroup := m.NewRouterGroup("/user/", s.mux, m.WithSessionManager(sm))
	userGroup.Post("login/", user.NewLogin(registry))
	userGroup.Get("logout/", user.NewLogout(registry), m.CheckLoginStatus)
	userGroup.Post("sign_up/", user.NewSignup(registry))
	userGroup.Get("get_user_info/", user.NewGetUserInfo(registry), m.AuthenticateUser)
	userGroup.Get("get_user_profile_picture/", user.NewGetUserProfilePicture(registry), m.CheckLoginStatus)
}

func (s *Server) Listen(host, port string) {
	log.Println("Listening on " + host + ":" + port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), s.mux)
	if err != nil {
		panic(err)
	}
}
