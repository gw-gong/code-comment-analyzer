package middleware

import (
	"net/http"

	"code-comment-analyzer/server/jwt"
)

func CheckLoginStatus(handlerFunc HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, extractor Extractor) {
		loginStatus := true
		userID, err := jwt.ParseToken(r, sessionManager)
		if err != nil {
			loginStatus = false
		}
		_ = jwt.RefreshToken(userID, sessionManager)
		if extractor == nil {
			extractor = newExtractedData()
		}
		extractor.setLoginStatus(loginStatus)
		extractor.setUserId(userID)
		handlerFunc(w, r, extractor)
	}
}
