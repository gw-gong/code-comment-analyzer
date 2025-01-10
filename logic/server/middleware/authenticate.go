package middleware

import (
	"code-comment-analyzer/data/redis"
	"net/http"

	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/jwt"
)

const CtxKeyUserID = "userIDFromAuthenticateForUser"

var sessionManager redis.SessionManager

func RegisterSessionManager(s redis.SessionManager) {
	sessionManager = s
}

func AuthenticateUser(handlerFunc HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, extractor Extractor) {
		userID, err := jwt.ParseToken(r, sessionManager)
		if err != nil {
			protocol.HandleError(w, protocol.ErrorCodeAuthenticating, err)
			return
		}
		err = jwt.RefreshToken(r, sessionManager)
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
