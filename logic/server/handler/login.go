package handler

import (
	"code-comment-analyzer/protocol"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"code-comment-analyzer/data"
	"code-comment-analyzer/server/jwt"
	"code-comment-analyzer/server/middleware"
)

type Login struct {
	w        http.ResponseWriter
	r        *http.Request
	registry *data.DataManagerRegistry
}

func NewLogin(registry *data.DataManagerRegistry) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &Login{
			w:        w,
			r:        r,
			registry: registry,
		}
	}
}

func (l *Login) Handle() {
	// 解析请求体中的 JSON 数据
	var requestData protocol.LoginRequest

	// 解码请求体的 JSON 数据到 requestData 结构体
	err := json.NewDecoder(l.r.Body).Decode(&requestData)
	if err != nil {
		log.Println("Failed to parse JSON body:", err)
		protocol.HttpResponseFail(l.w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	// 检查 email 和 password 是否为空
	if requestData.Email == "" || requestData.Password == "" {
		log.Println("Email or password is missing")
		// 如果 email 或 password 为空，返回错误响应
		protocol.HttpResponseFail(l.w, http.StatusBadRequest, "Email or password is missing")
		return
	}

	um := l.registry.GetUserManager()
	userID, nickname, password, err := um.GetUserInfoByEmail(requestData.Email)
	if err != nil {
		log.Printf("Error|GetUserInfoByEmail|err: %v", err)
		protocol.HttpResponseFail(l.w, http.StatusBadRequest, fmt.Sprintf("Error|GetUserInfoByEmail|err: %v", err))
	}

	// 假设你会根据 email 和 password 验证用户（可以在数据库中查询或其他验证方式）
	// 这里只是简单模拟用户验证
	if requestData.Password != password {
		log.Println("Invalid email or password")
		// 如果验证失败，返回错误响应
		protocol.HttpResponseFail(l.w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// 用户验证成功，生成 JWT Token
	err = jwt.AuthorizeUserToken(userID, l.w, l.registry.GetSessionManager())
	if err != nil {
		// 如果授权失败，返回错误响应
		protocol.HttpResponseFail(l.w, http.StatusUnauthorized, "Authorization failed")
		return
	}

	// 假设通过授权后可以获取用户的其他信息（如 email, nickname），在此简单模拟：
	// 获取用户信息（例如：通过 userID 查询数据库）
	// 返回成功响应
	response := protocol.LoginResponse{
		UID:      userID,
		Email:    requestData.Email,
		Nickname: nickname,
	}
	protocol.HttpResponseSuccess(l.w, http.StatusOK, "登录成功", response)
}
