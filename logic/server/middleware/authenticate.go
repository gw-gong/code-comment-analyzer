package middleware

import (
	"net/http"

	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/jwt"
)

const CtxKeyUserID = "userIDFromAuthenticateForUser"

func AuthenticateForUser(handlerFunc HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, extractor Extractor) {
		userID, err := jwt.ParseToken(r)
		if err != nil {
			protocol.HandleError(w, protocol.ErrorCodeAuthenticating, err)
			return
		}
		if extractor == nil {
			extractor = newExtractedData()
		}
		extractor.setUserId(userID)
		handlerFunc(w, r, extractor)
	}
}
