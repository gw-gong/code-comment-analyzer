package user

import (
	"net/http"

	"code-comment-analyzer/server/middleware"
)

type UserPageRedirect struct {
	w         http.ResponseWriter
	r         *http.Request
	extractor middleware.Extractor
}

func NewUserPageRedirect() middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &UserPageRedirect{
			w:         w,
			r:         r,
			extractor: extractor,
		}
	}
}

func (u *UserPageRedirect) Handle() {
	if isUserLoggedIn, err := u.extractor.IsUserLoggedIn(); err != nil || !isUserLoggedIn {
		http.Redirect(u.w, u.r, "/login/", http.StatusMovedPermanently)
		return
	}
	http.Redirect(u.w, u.r, "/user_info/", http.StatusMovedPermanently)
}
