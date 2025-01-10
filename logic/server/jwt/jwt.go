package jwt

import (
	"code-comment-analyzer/config"
	"code-comment-analyzer/data/redis"
	"code-comment-analyzer/protocol"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func AuthorizeUserToken(userID uint64, w http.ResponseWriter, sessionManager redis.SessionManager) {
	jwtKey := []byte(config.Cfg.JwtKey)
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		protocol.HandleError(w, protocol.ErrorCodeAuthorizing, err)
		log.Println(err)
		return
	}
	err = sessionManager.SetSession(tokenString, userID)
	if err != nil {
		protocol.HandleError(w, protocol.ErrorCodeAuthorizing, err)
		log.Println(err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: tokenString,
	})
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Login successful and token has been set."))
	if err != nil {
		protocol.HandleError(w, protocol.ErrorCodeAuthorizing, err)
		return
	}
}

func ParseToken(r *http.Request, sessionManager redis.SessionManager) (userID uint64, err error) {
	c, err := r.Cookie("token")
	if err != nil {
		return 0, err
	}
	tokenString := c.Value
	jwtKey := []byte(config.Cfg.JwtKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}
	return sessionManager.GetSession(tokenString)
}

func RefreshToken(r *http.Request, sessionManager redis.SessionManager) error {
	c, err := r.Cookie("token")
	if err != nil {
		return err
	}
	tokenString := c.Value
	err = sessionManager.RefreshSession(tokenString)
	if err != nil {
		return err
	}
	return nil
}
