package user

import (
	"fmt"
	"net/http"

	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
	"code-comment-analyzer/util"
)

type GetUserInfo struct {
	w         http.ResponseWriter
	r         *http.Request
	extractor middleware.Extractor
	registry  *data.DataManagerRegistry
}

func NewGetUserInfo(registry *data.DataManagerRegistry) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &GetUserInfo{
			w:         w,
			r:         r,
			extractor: extractor,
			registry:  registry,
		}
	}
}

func (g *GetUserInfo) Handle() {
	userID, err := g.extractor.GetUserId()
	if err != nil {
		protocol.HttpResponseFail(g.w, http.StatusInternalServerError, protocol.ErrorCodeMissingUserId, fmt.Sprintf("%v", err))
		return
	}

	um := g.registry.GetUserManager()
	isSetProfilePicture, profilePictureUrl, email, nickname, dateJoined, err := um.GetUserInfoByUserID(userID)
	if err != nil {
		protocol.HttpResponseFail(g.w, http.StatusInternalServerError, protocol.ErrorCodeInternalServerError, fmt.Sprintf("%v", err))
		return
	}

	tableInfo := protocol.TableInfo{
		Email:      email,
		NickName:   nickname,
		DateJoined: dateJoined.Format("2006-01-02"),
	}

	response := &protocol.GetUserInfoResponse{
		ProfilePicture: *(util.TransformProfilePictureUrlToResourceUrl(isSetProfilePicture, profilePictureUrl)),
		TableInfo:      []protocol.TableInfo{tableInfo},
	}

	protocol.HttpResponseSuccess(g.w, http.StatusOK, "获取用户信息成功", protocol.WithData(response))
}
