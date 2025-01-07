package middleware

import (
	"fmt"
	"net/http"
)

func AuthenticationForUser(handlerFunc HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 鉴权逻辑
		// 从 Cookie 中提取 Token
		cookie, err := r.Cookie("AuthToken")
		if err != nil {
			// 如果没有 Token 或无法提取，返回错误
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			fmt.Println("Authentication failed: No token provided")
			return
		}

		// TODO: 在这里添加更详细的 Token 验证逻辑
		token := cookie.Value
		fmt.Println("Token received:", token)

		// 假设我们接受 "xxxxxx" 作为有效 Token
		if token != "4xxxxx" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			fmt.Println("Authentication failed: Invalid token")
			return
		}

		fmt.Println("Authentication successful")
		// Token 验证成功，调用原始的处理函数

		fmt.Println("Authentication For User")
		handlerFunc(w, r)
	}
}

// 使用jwt
