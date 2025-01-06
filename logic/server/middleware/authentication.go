package middleware

import (
	"fmt"
	"net/http"
)

func AuthenticationForUser(handlerFunc HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 鉴权逻辑
		fmt.Println("Authentication For User")
		handlerFunc(w, r)
	}
}

// 使用jwt
