package jwt

import (
	"fmt"
	"net/http"
	"time"

	"code-comment-analyzer/config"
	"code-comment-analyzer/protocol"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID uint64 `json:"user_id"`
	jwt.StandardClaims
}

func AuthorizeUserToken(userID uint64, w http.ResponseWriter) {
	duration := time.Duration(config.Cfg.UserTokenDuration)
	jwtKey := []byte(config.Cfg.JwtKey)
	expireTime := time.Now().Add(duration * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		protocol.HandleError(w, protocol.ErrorCodeAuthorizing, err)
		fmt.Println(err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expireTime,
	})
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Login successful and token has been set."))
	if err != nil {
		protocol.HandleError(w, protocol.ErrorCodeAuthorizing, err)
		return
	}
}

func ParseToken(r *http.Request) (userID uint64, err error) {
	c, err := r.Cookie("token")
	if err != nil {
		return 0, err
	}
	tokenString := c.Value
	claims := &Claims{}
	jwtKey := []byte(config.Cfg.JwtKey)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, fmt.Errorf(protocol.ErrorMsgInvalidToken)
	}
	return claims.UserID, nil
}
