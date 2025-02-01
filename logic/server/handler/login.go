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
	// 解析请求体中的 JSON 数据
	var requestData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// 解码请求体的 JSON 数据到 requestData 结构体
	err := json.NewDecoder(l.r.Body).Decode(&requestData)
	if err != nil {
		log.Println("Failed to parse JSON body:", err)
		// 如果解析失败，返回错误响应
		response := map[string]interface{}{
			"status": 1,
			"msg":    "Invalid JSON format",
			"data":   map[string]interface{}{},
		}
		l.w.Header().Set("Content-Type", "application/json")
		l.w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(l.w).Encode(response)
		return
	}

	// 检查 email 和 password 是否为空
	if requestData.Email == "" || requestData.Password == "" {
		log.Println("Email or password is missing")
		// 如果 email 或 password 为空，返回错误响应
		response := map[string]interface{}{
			"status": 1,
			"msg":    "Email or password is missing",
			"data":   map[string]interface{}{},
		}
		l.w.Header().Set("Content-Type", "application/json")
		l.w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(l.w).Encode(response)
		return
	}

	// 假设你会根据 email 和 password 验证用户（可以在数据库中查询或其他验证方式）
	// 这里只是简单模拟用户验证
	if requestData.Email != "xxx@qq.com" || requestData.Password != "123456" {
		log.Println("Invalid email or password")
		// 如果验证失败，返回错误响应
		response := map[string]interface{}{
			"status": 1,
			"msg":    "Invalid email or password",
			"data":   map[string]interface{}{},
		}
		l.w.Header().Set("Content-Type", "application/json")
		l.w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(l.w).Encode(response)
		return
	}

	// 用户验证成功，生成 JWT Token
	userID := uint64(3) // 模拟用户ID，实际应该从数据库获取
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
		json.NewEncoder(l.w).Encode(response)
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
		Email:    requestData.Email,
		Nickname: "xxxxxxxxxxxxx", // 示例：模拟的昵称
	}

	// 返回成功响应
	response := map[string]interface{}{
		"status": 0,
		"msg":    "登录成功",
		"data":   user,
	}

	l.w.Header().Set("Content-Type", "application/json")
	l.w.WriteHeader(http.StatusOK)
	json.NewEncoder(l.w).Encode(response)
}
