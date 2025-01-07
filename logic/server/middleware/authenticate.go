package middleware

import (
	"context"
	"net/http"

	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/jwt"
)

const CtxKeyUserID = "userIDFromAuthenticateForUser"

func AuthenticateForUser(handlerFunc HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := jwt.ParseToken(r)
		if err != nil {
			protocol.HandleError(w, protocol.ErrorCodeAuthenticating, err)
			return
		}
		ctx := context.WithValue(r.Context(), CtxKeyUserID, userID)
		rWithUser := r.WithContext(ctx)
		handlerFunc(w, rWithUser)
	}
}
