package jwt

import (
	"fmt"
	"log"
	"net/http"

	"code-comment-analyzer/config"
	"code-comment-analyzer/data/redis"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	jwt.StandardClaims
	UserID uint64 `json:"user_id"`
}

func AuthorizeUserToken(userID uint64, w http.ResponseWriter, sessionManager redis.SessionManager) error {
	userClaims := &UserClaims{
		UserID: userID,
	}
	jwtKey := []byte(config.Cfg.JwtKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error signing token: %w", err)
	}
	err = sessionManager.SetSession(userID, tokenString)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error setting session: %w", err)
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: tokenString,
	})
	return nil
}

func ParseToken(r *http.Request, sessionManager redis.SessionManager) (userID uint64, err error) {
	c, err := r.Cookie("token")
	if err != nil {
		return 0, err
	}
	tokenString := c.Value
	jwtKey := []byte(config.Cfg.JwtKey)
	userClaims := &UserClaims{}
	token, err := jwt.ParseWithClaims(tokenString, userClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}
	validToken, err := sessionManager.GetSession(userClaims.UserID)
	if err != nil {
		return 0, err
	}
	if tokenString != validToken {
		return 0, fmt.Errorf("invalid token")
	}
	return userClaims.UserID, nil
}

func RefreshToken(userID uint64, sessionManager redis.SessionManager) error {
	err := sessionManager.RefreshSession(userID)
	if err != nil {
		return err
	}
	return nil
}
