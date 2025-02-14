package user

import (
	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
	"code-comment-analyzer/util"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

type UpdateUserInfo struct {
	w         http.ResponseWriter
	r         *http.Request
	extractor middleware.Extractor
	registry  *data.DataManagerRegistry
}



func NewUpdateUserInfo(registry *data.DataManagerRegistry) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &UpdateUserInfo{
			w:         w,
			r:         r,
			extractor: extractor,
			registry:  registry,
		}
	}
}

func (u *UpdateUserInfo) Handle() {
	// 获取用户ID
	userID, err := u.extractor.GetUserId()
	if err != nil {
		protocol.HttpResponseFail(u.w, http.StatusInternalServerError, protocol.ErrorCodeMissingUserId, fmt.Sprintf("%v", err))
		return
	}

	// 解析multipart表单
	if err := u.r.ParseMultipartForm(10 << 20); err != nil { // 10 MB 限制
		protocol.HttpResponseFail(u.w, http.StatusBadRequest, protocol.ErrorCodeInvalidRequest, "无法解析表单数据")
		return
	}

	// 获取用户信息管理器
	userManager := u.registry.GetUserManager()

	// 处理头像上传
	file, header, err := u.r.FormFile("profile_picture")
	if err == nil && file != nil {
		defer file.Close()

		// 生成唯一文件名
		avatarFileName := uuid.New().String() + filepath.Ext(header.Filename)

		// 获取存储路径
		storagePath := util.GetAvatarStoragePath(userID, avatarFileName)

		// 确保目录存在
		if err := os.MkdirAll(filepath.Dir(storagePath), 0755); err != nil {
			protocol.HttpResponseFail(u.w, http.StatusInternalServerError, protocol.ErrorCodeInternalError, "创建存储目录失败")
			return
		}

		// 创建目标文件
		dst, err := os.Create(storagePath)
		if err != nil {
			protocol.HttpResponseFail(u.w, http.StatusInternalServerError, protocol.ErrorCodeInternalError, "创建文件失败")
			return
		}
		defer dst.Close()

		// 保存文件
		if _, err := io.Copy(dst, file); err != nil {
			protocol.HttpResponseFail(u.w, http.StatusInternalServerError, protocol.ErrorCodeInternalError, "保存头像失败")
			return
		}

		// 更新数据库中的头像信息
		if err := userManager.UpdateUserAvatar(userID, avatarFileName); err != nil {
			protocol.HttpResponseFail(u.w, http.StatusInternalServerError, protocol.ErrorCodeInternalError, "更新头像信息失败")
			return
		}
	}

	var updateInfo protocol.UpdateUserInfoRequest
	if err := json.NewDecoder(strings.NewReader(u.r.FormValue("json"))).Decode(&updateInfo); err != nil {
		protocol.HttpResponseFail(u.w, http.StatusBadRequest, protocol.ErrorCodeInvalidRequest, "无效的JSON数据")
		return
	}

	// 更新用户信息
	if err := userManager.UpdateUserInfo(userID, updateInfo.Nickname, updateInfo.AgainNewPassword); err != nil {
		protocol.HttpResponseFail(u.w, http.StatusInternalServerError, protocol.ErrorCodeInternalError, "更新用户信息失败")
		return
	}

	protocol.HttpResponseSuccess(u.w, http.StatusOK, "用户信息更新成功", nil)
}
