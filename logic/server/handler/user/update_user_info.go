package user

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"code-comment-analyzer/config"
	"code-comment-analyzer/data"
	"code-comment-analyzer/data/mysql"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
	"code-comment-analyzer/util"
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
	userID, err := u.extractor.GetUserId()
	if err != nil {
		protocol.HttpResponseFail(u.w, http.StatusInternalServerError, protocol.ErrorCodeMissingUserId, fmt.Sprintf("%v", err))
		return
	}

	if err := u.r.ParseMultipartForm(config.Cfg.MaxAvatarSize << 20); err != nil {
		protocol.HttpResponseFail(u.w, http.StatusBadRequest, protocol.ErrorCodeInvalidRequest, "无法解析表单数据")
		return
	}

	userManager := u.registry.GetUserManager()

	if err := u.handleProfilePicture(userID, userManager); err != nil {
		protocol.HttpResponseFail(u.w, http.StatusInternalServerError, protocol.ErrorCodeInternalError, err.Error())
		return
	}

	if err := u.updateUserInfo(userID, userManager); err != nil {
		protocol.HttpResponseFail(u.w, http.StatusInternalServerError, protocol.ErrorCodeInternalError, err.Error())
		return
	}

	protocol.HttpResponseSuccess(u.w, http.StatusOK, "用户信息更新成功")
}

func (u *UpdateUserInfo) handleProfilePicture(userID uint64, userManager mysql.UserManager) error {
	file, header, err := u.r.FormFile("profile_picture")
	if err != nil || file == nil {
		return nil
	}
	defer file.Close()

	avatarFileName := util.GenerateUUIDAvatarName() + filepath.Ext(header.Filename)
	storagePath := util.GetAvatarStoragePath(userID, avatarFileName)

	if err := os.MkdirAll(filepath.Dir(storagePath), 0755); err != nil {
		return fmt.Errorf("创建存储目录失败")
	}

	dst, err := os.Create(storagePath)
	if err != nil {
		return fmt.Errorf("创建文件失败")
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return fmt.Errorf("保存头像失败")
	}

	if err := userManager.UpdateUserAvatar(userID, avatarFileName); err != nil {
		return fmt.Errorf("更新头像信息失败")
	}

	return nil
}

func (u *UpdateUserInfo) updateUserInfo(userID uint64, userManager mysql.UserManager) error {
	var updateInfo protocol.UpdateUserInfoRequest
	if err := json.NewDecoder(strings.NewReader(u.r.FormValue("json"))).Decode(&updateInfo); err != nil {
		return fmt.Errorf("无效的JSON数据")
	}

	if err := userManager.UpdateUserInfo(userID, updateInfo.Nickname, updateInfo.AgainNewPassword); err != nil {
		return fmt.Errorf("更新用户信息失败")
	}

	return nil
}
