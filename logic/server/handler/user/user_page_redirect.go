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
		http.Redirect(u.w, u.r, "/login/", http.StatusFound)
		return
	}
	http.Redirect(u.w, u.r, "/user_info/", http.StatusFound)
}

// 注意：
// 使用http.StatusMovedPermanently，表示永久重定向，即浏览器会缓存该重定向。缓存之后直接不走后台，直接重定向到缓存的地址。不符合预期
// 使用http.StatusFound，表示临时重定向，即浏览器不会缓存该重定向。
