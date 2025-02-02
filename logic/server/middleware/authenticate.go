package middleware

import (
	"fmt"
	"net/http"

	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/jwt"
)

func AuthenticateUser(handlerFunc HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, extractor Extractor) {
		userID, err := jwt.ParseToken(r, sessionManager)
		if err != nil {
			protocol.HttpResponseFail(w, http.StatusInternalServerError, protocol.ErrorCodeAuthenticating, fmt.Sprintf("%v", err))
			return
		}
		err = jwt.RefreshToken(userID, sessionManager)
		if err != nil {
			protocol.HttpResponseFail(w, http.StatusInternalServerError, protocol.ErrorCodeAuthenticating, fmt.Sprintf("%v", err))
			return
		}
		if extractor == nil {
			extractor = newExtractedData()
		}
		extractor.setUserId(userID)
		handlerFunc(w, r, extractor)
	}
}
