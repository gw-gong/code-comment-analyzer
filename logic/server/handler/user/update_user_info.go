package user

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

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

	avatarWantUpdate := false
	infoWantUpdate := false
	avatarErrChan := make(chan error, 1)
	infoErrChan := make(chan error, 1)

	go util.WithRecover(func() {
		avatarWantUpdate, err = u.handleProfilePicture(userID, userManager)
		avatarErrChan <- err
	})

	go util.WithRecover(func() {
		infoWantUpdate, err = u.updateUserInfo(userID, userManager)
		infoErrChan <- err
	})

	avatarErr := <-avatarErrChan
	infoErr := <-infoErrChan

	if avatarWantUpdate && infoWantUpdate {
		if avatarErr != nil && infoErr == nil {
			protocol.HttpResponseSuccess(u.w, http.StatusOK, "更新昵称成功，但更新头像失败")
		} else if avatarErr == nil && infoErr != nil {
			protocol.HttpResponseSuccess(u.w, http.StatusOK, "更新头像成功，但更新昵称失败")
		} else if avatarErr != nil && infoErr != nil {
			protocol.HttpResponseFail(u.w, http.StatusInternalServerError, protocol.ErrorCodeInternalError, "用户信息更新失败")
		} else {
			protocol.HttpResponseSuccess(u.w, http.StatusOK, "用户信息更新成功")
		}
		return
	}

	if avatarWantUpdate && avatarErr != nil {
		protocol.HttpResponseFail(u.w, http.StatusInternalServerError, protocol.ErrorCodeInternalError, "更新头像失败")
		return
	}

	if infoWantUpdate && infoErr != nil {
		protocol.HttpResponseFail(u.w, http.StatusInternalServerError, protocol.ErrorCodeInternalError, "更新昵称失败")
		return
	}

	protocol.HttpResponseSuccess(u.w, http.StatusOK, "用户信息更新成功")
}

func (u *UpdateUserInfo) handleProfilePicture(userID uint64, userManager mysql.UserManager) (wantUpdate bool, err error) {
	file, header, err := u.r.FormFile(protocol.MultipartFormKeyProfilePicture)
	if err != nil || file == nil {
		// 这里是因为只要有上传就更新，没填写就不管
		return false, nil
	}
	defer file.Close()

	avatarFileName := util.GenerateUUIDAvatarName() + filepath.Ext(header.Filename)
	storagePath := util.GetAvatarStoragePath(userID, avatarFileName)

	if err := os.MkdirAll(filepath.Dir(storagePath), 0755); err != nil {
		return true, fmt.Errorf("创建存储目录失败")
	}

	dst, err := os.Create(storagePath)
	if err != nil {
		return true, fmt.Errorf("创建文件失败")
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return true, fmt.Errorf("保存头像失败")
	}

	if err := userManager.UpdateUserAvatar(userID, avatarFileName); err != nil {
		return true, fmt.Errorf("更新头像信息失败")
	}

	return true, nil
}

func (u *UpdateUserInfo) updateUserInfo(userID uint64, userManager mysql.UserManager) (wantUpdate bool, err error) {
	nickname := u.r.FormValue(protocol.FormKeyNickname)

	// 这里是因为只要有填写就更新，没填写就不管
    if nickname == "" {
        return false, nil  // 如果没有提供昵称，不进行更新
    }

	if err := userManager.UpdateUserInfo(userID, nickname); err != nil {
		return true, fmt.Errorf("更新用户昵称失败")
	}

	return true, nil
}
