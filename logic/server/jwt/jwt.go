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
	// 创建用户的声明信息
	userClaims := &UserClaims{
		UserID: userID,
	}

	// 生成 JWT 密钥
	jwtKey := []byte(config.Cfg.JwtKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	// 签名生成 token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// 如果签名失败，返回错误
		log.Println(err)
		return fmt.Errorf("error signing token: %w", err)
	}

	// 将 token 保存到会话中
	err = sessionManager.SetSession(userID, tokenString)
	if err != nil {
		// 如果保存会话失败，返回错误
		log.Println(err)
		return fmt.Errorf("error setting session: %w", err)
	}

	// 设置响应中的 token Cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: tokenString,
	})

	// 返回成功响应
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Login successful and token has been set."))
	if err != nil {
		// 如果写入响应体失败，返回错误
		log.Println(err)
		return fmt.Errorf("error writing response body: %w", err)
	}

	// 如果没有发生错误，返回 nil
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
