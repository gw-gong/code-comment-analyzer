package middleware

import (
	"fmt"
	"net/http"

	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/jwt"
)

func AuthenticateForUser(handlerFunc HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := jwt.ParseToken(r)
		if err != nil {
			protocol.HandleError(w, protocol.ErrorCodeAuthenticating, err)
			return
		}
		// 待处理如何传递userID
		fmt.Println(userID)
		handlerFunc(w, r)
	}
}

// 使用jwt
