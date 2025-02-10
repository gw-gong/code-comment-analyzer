package middleware

import (
	"net/http"

	"code-comment-analyzer/server/jwt"
)

func CheckLoginStatus(rg *routerGroup, handlerFunc HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, extractor Extractor) {
		loginStatus := true
		userID, err := jwt.ParseToken(r, rg.getSessionManager())
		if err != nil {
			loginStatus = false
		}
		_ = jwt.RefreshToken(userID, rg.getSessionManager())
		if extractor == nil {
			extractor = newExtractedData()
		}
		extractor.setLoginStatus(loginStatus)
		extractor.setUserId(userID)
		handlerFunc(w, r, extractor)
	}
}
