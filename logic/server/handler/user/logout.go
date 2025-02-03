package user

import (
	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
	"log"
	"net/http"
)

type Logout struct {
	w        http.ResponseWriter
	r        *http.Request
	registry *data.DataManagerRegistry
}

func NewLogout(registry *data.DataManagerRegistry) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &Logout{
			w:        w,
			r:        r,
			registry: registry,
		}
	}
}

func (l *Logout) Handle() {
	// 检查请求中的 token（通过 Cookie 获取）
	tokenCookie, err := l.r.Cookie("token")
	if err != nil || tokenCookie == nil {
		log.Println("No token found in request")
		// 如果没有 token，认为用户已经退出（无 token 也返回成功）
		protocol.HttpResponseSuccess(l.w, http.StatusOK, "用户已成功退出登录", map[string]interface{}{})
		return
	}

	// 删除浏览器中的 token Cookie
	http.SetCookie(l.w, &http.Cookie{
		Name:   "token",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // 使 Cookie 立即过期
	})

	// 返回成功响应
	protocol.HttpResponseSuccess(l.w, http.StatusOK, "已成功退出登录", map[string]interface{}{})
}
