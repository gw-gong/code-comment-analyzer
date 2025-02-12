package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
)

type ChangePassword struct {
	w         http.ResponseWriter
	r         *http.Request
	extractor middleware.Extractor
	registry  *data.DataManagerRegistry
}

func NewChangePassword(registry *data.DataManagerRegistry) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &ChangePassword{
			w:         w,
			r:         r,
			extractor: extractor,
			registry:  registry,
		}
	}
}

func (c *ChangePassword) Handle() {
	// 解析请求中的 JSON 数据
	ChangePasswordRequest, err := c.decodeRequest()
	if err != nil {
		protocol.HttpResponseFail(c.w, http.StatusBadRequest, protocol.ErrorCodeInvalidRequest, "请求数据无效")
		return
	}

	// 校验密码是否一致
	if ChangePasswordRequest.NewPassword != ChangePasswordRequest.AgainNewPassword {
		protocol.HttpResponseFail(c.w, http.StatusBadRequest, protocol.ErrorCodeInvalidPassword, "两次输入的新密码不一致")
		return
	}

	// 获取当前用户 ID
	userID, err := c.extractor.GetUserId()
	if err != nil {
		protocol.HttpResponseFail(c.w, http.StatusInternalServerError, protocol.ErrorCodeMissingUserId, fmt.Sprintf("%v", err))
		return
	}

	// 获取用户管理器
	um := c.registry.GetUserManager()

	// 验证旧密码是否正确 以及更新密码
	err = um.CheckOldPassword(userID, ChangePasswordRequest.OldPassword, ChangePasswordRequest.NewPassword)
	if err != nil {
		protocol.HttpResponseFail(c.w, http.StatusUnauthorized, protocol.ErrorCodeInvalidPassword, "旧密码错误")
		return
	}
	// 返回成功响应
	protocol.HttpResponseSuccess(c.w, http.StatusOK, "密码更新成功", protocol.WithData(map[string]interface{}{}))
}

func (c *ChangePassword) decodeRequest() (*protocol.ChangePasswordRequest, error) {
	var requestData protocol.ChangePasswordRequest
	// 解码请求体的 JSON 数据到 requestData 结构体
	err := json.NewDecoder(c.r.Body).Decode(&requestData)
	if err != nil {
		log.Println("Failed to parse JSON body:", err)
		protocol.HttpResponseFail(c.w, http.StatusBadRequest, protocol.ErrorCodeParseRequestFailed, "Invalid JSON format")
		return nil, err
	}
	return &requestData, nil
}
