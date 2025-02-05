package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"code-comment-analyzer/config"
	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
)

type Signup struct {
	w        http.ResponseWriter
	r        *http.Request
	registry *data.DataManagerRegistry
}

func NewSignup(registry *data.DataManagerRegistry) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &Signup{
			w:        w,
			r:        r,
			registry: registry,
		}
	}
}

func (s *Signup) Handle() {
	// 解码请求数据
	requestData, err := s.decodeRequest()
	if err != nil {
		return
	}

	// 检查 email 和 password 是否为空
	if requestData.Email == "" || requestData.Password == "" || requestData.PasswordAgain == "" {
		log.Println("Email or password is missing")
		protocol.HttpResponseFail(s.w, http.StatusBadRequest, protocol.ErrorCodeParamError, "Email or password is missing")
		return
	}

	// 检查用户输入两边密码是否一致
	if requestData.Password != requestData.PasswordAgain {
		log.Println("Passwords do not match")
		protocol.HttpResponseFail(s.w, http.StatusBadRequest, protocol.ErrorCodeParamError, "Passwords do not match")
		return
	}

	// 检查邮箱是否已被注册
	um := s.registry.GetUserManager()
	isExistUser, err := um.IsExistUserByEmail(requestData.Email)
	if err != nil {
		log.Printf("Error|GetUserByEmail|err: %v", err)
		protocol.HttpResponseFail(s.w, http.StatusBadRequest, protocol.ErrorCodeRegisteredEmail, fmt.Sprintf("%v", err))
		return
	}
	if isExistUser {
		log.Println("Email already registered")
		protocol.HttpResponseFail(s.w, http.StatusBadRequest, protocol.ErrorCodeRegisteredEmail, "Email already registered")
		return
	}

	// 设置默认的nickname
	defaultNickname := config.Cfg.DefaultNickname

	// 创建新用户并保存
	userID, err := um.CreateUser(requestData.Email, requestData.Password, defaultNickname)
	if err != nil {
		log.Printf("Error|CreateUser|err: %v", err)
		protocol.HttpResponseFail(s.w, http.StatusInternalServerError, protocol.ErrorCodeInternalServerError, fmt.Sprintf("Error|CreateUser|err: %v", err))
		return
	}

	// 返回注册成功的响应
	response := &protocol.SignupResponse{
		UID:      userID,
		Email:    requestData.Email,
		Nickname: defaultNickname,
	}
	protocol.HttpResponseSuccess(s.w, http.StatusOK, "注册成功", response)
}

func (s *Signup) decodeRequest() (*protocol.SignupRequest, error) {
	var requestData protocol.SignupRequest
	// 解码请求体的 JSON 数据到 requestData 结构体
	err := json.NewDecoder(s.r.Body).Decode(&requestData)
	if err != nil {
		log.Println("Failed to parse JSON body:", err)
		protocol.HttpResponseFail(s.w, http.StatusBadRequest, protocol.ErrorCodeParseRequestFailed, "Invalid JSON format")
		return nil, err
	}
	return &requestData, nil
}
