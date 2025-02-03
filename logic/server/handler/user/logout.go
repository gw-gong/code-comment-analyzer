package user

import (
	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
	"fmt"
	"net/http"
)

type Logout struct {
	w         http.ResponseWriter
	r         *http.Request
	extractor middleware.Extractor
	registry  *data.DataManagerRegistry
}

func NewLogout(registry *data.DataManagerRegistry) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &Logout{
			w:         w,
			r:         r,
			extractor: extractor,
			registry:  registry,
		}
	}
}

func (l *Logout) Handle() {
	// 获取登录状态，未登录直接成功
	loginStatus, err := l.extractor.GetLoginStatus()
	if err != nil {
		protocol.HttpResponseFail(l.w, http.StatusBadRequest, protocol.ErrorCodeInternalServerError, fmt.Sprintf("%v", err))
		return
	}
	if !loginStatus {
		protocol.HttpResponseSuccess(l.w, http.StatusOK, "success! 用户本来未登录", nil)
		return
	}
	// 从token中获取UserID，使用UserID清除redis中的token，ToDo: 是否有必要清除浏览器的token
	userID, err := l.extractor.GetUserId()
	if err != nil {
		protocol.HttpResponseFail(l.w, http.StatusInternalServerError, protocol.ErrorCodeMissingUserId, fmt.Sprintf("%v", err))
		return
	}
	sm := l.registry.GetSessionManager()
	err = sm.ClearSession(userID)
	if err != nil {
		protocol.HttpResponseFail(l.w, http.StatusInternalServerError, protocol.ErrorCodeInternalServerError, fmt.Sprintf("%v", err))
		return
	}

	protocol.HttpResponseSuccess(l.w, http.StatusOK, "已成功退出登录", nil)
}
