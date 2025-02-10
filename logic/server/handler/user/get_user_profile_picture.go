package user

import (
	"fmt"
	"net/http"
	"path/filepath"

	"code-comment-analyzer/config"
	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
)

type GetUserProfilePicture struct {
	w         http.ResponseWriter
	r         *http.Request
	extractor middleware.Extractor
	registry  *data.DataManagerRegistry
}

func NewGetUserProfilePicture(registry *data.DataManagerRegistry) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &GetUserProfilePicture{
			w:         w,
			r:         r,
			extractor: extractor,
			registry:  registry,
		}
	}
}

func (g *GetUserProfilePicture) Handle() {
	isUserLoggedIn, err := g.extractor.IsUserLoggedIn()
	if err != nil {
		protocol.HttpResponseFail(g.w, http.StatusBadRequest, protocol.ErrorCodeInternalServerError, fmt.Sprintf("%v", err))
		return
	}

	response := &protocol.GetUserProfilePictureResponse{}

	if !isUserLoggedIn {
		response.Text = "?"
		protocol.HttpResponseSuccess(g.w, http.StatusOK, "未登录", protocol.WithData(response))
		return
	}

	userID, err := g.extractor.GetUserId()
	if err != nil {
		protocol.HttpResponseFail(g.w, http.StatusInternalServerError, protocol.ErrorCodeMissingUserId, fmt.Sprintf("%v", err))
		return
	}

	um := g.registry.GetUserManager()
	isSetProfilePicture, profilePictureUrl, err := um.GetUserProfilePictureByUserID(userID)
	if err != nil {
		protocol.HttpResponseFail(g.w, http.StatusInternalServerError, protocol.ErrorCodeInternalServerError, fmt.Sprintf("%v", err))
		return
	}

	if isSetProfilePicture {
		// todo 和上传头像接口协商一下，以及获取头像需要鉴权，这块需要设计一下
		response.ProfilePicture = &profilePictureUrl
		// response.Text = ？？？看一下python项目代码逻辑
	} else {
		avatarsStorageRootPath := config.Cfg.FileStoragePath.Avatars
		defaultAvatar := config.Cfg.DefaultAvatar
		defaultAvatarUrl := filepath.Join(avatarsStorageRootPath, defaultAvatar)
		response.ProfilePicture = &defaultAvatarUrl
		response.Text = "未设置"
	}

	protocol.HttpResponseSuccess(g.w, http.StatusOK, "获取用户头像成功", protocol.WithData(response))
}
