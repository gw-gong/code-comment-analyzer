package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"code-comment-analyzer/data"
	"code-comment-analyzer/server/jwt"
	"code-comment-analyzer/server/middleware"
)

type Login struct {
	w         http.ResponseWriter
	r         *http.Request
	extractor middleware.Extractor
	registry  *data.DataManagerRegistry
}

func NewLogin(registry *data.DataManagerRegistry) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &Login{
			w:         w,
			r:         r,
			extractor: extractor,
			registry:  registry,
		}
	}
}
func (l *Login) Handle() {
	// 获取提取的数据（如 userID）
	if l.extractor == nil {
		log.Println("Extractor is nil")
		return
	}
	userID, err := l.extractor.GetUserId()
	if err != nil {
		// 如果没有找到用户ID，返回错误响应
		response := map[string]interface{}{
			"status": 1,
			"msg":    "User ID not found",
			"data":   map[string]interface{}{},
		}
		l.w.Header().Set("Content-Type", "application/json")
		l.w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(l.w).Encode(response) // 使用 json.NewEncoder 来编码响应
		return
	}

	// 调用 jwt.AuthorizeUserToken 来授权用户
	err = jwt.AuthorizeUserToken(userID, l.w, l.registry.GetSessionManager())
	if err != nil {
		// 如果授权失败，返回错误响应
		response := map[string]interface{}{
			"status": 1,
			"msg":    "Authorization failed",
			"data":   map[string]interface{}{},
		}
		l.w.Header().Set("Content-Type", "application/json")
		l.w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(l.w).Encode(response) // 使用 json.NewEncoder 来编码响应
		return
	}

	// 假设通过授权后可以获取用户的其他信息（如 email, nickname），在此简单模拟：
	// 获取用户信息（例如：通过 userID 查询数据库）
	user := struct {
		UID      uint64 `json:"uid"`
		Email    string `json:"email"`
		Nickname string `json:"nickname"`
	}{
		UID:      userID,
		Email:    "paipai@qq.com", // 示例：模拟的邮箱
		Nickname: "xxxxxxxxxxxxx", // 示例：模拟的昵称
	}

	// 如果授权成功，返回成功响应
	response := map[string]interface{}{
		"status": 0,
		"msg":    "登录成功",
		"data":   user,
	}

	l.w.Header().Set("Content-Type", "application/json")
	l.w.WriteHeader(http.StatusOK)
	json.NewEncoder(l.w).Encode(response) // 使用 json.NewEncoder 来编码响应
}
