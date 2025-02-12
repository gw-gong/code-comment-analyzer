package file_storage

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"code-comment-analyzer/config"
	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
)

type GetAvatars struct {
	w         http.ResponseWriter
	r         *http.Request
	extractor middleware.Extractor
	registry  *data.DataManagerRegistry
}

func NewGetAvatars(registry *data.DataManagerRegistry) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &GetAvatars{
			w:         w,
			r:         r,
			extractor: extractor,
			registry:  registry,
		}
	}
}

func (g *GetAvatars) Handle() {
	avatarPath := strings.TrimLeft(g.r.URL.Path, "/")
	avatarFileName := filepath.Base(avatarPath)
	avatarDirPath := filepath.Dir(avatarPath)
	if avatarFileName == config.Cfg.DefaultAvatar {
		if _, err := os.Stat(avatarPath); os.IsNotExist(err) {
			protocol.HttpResponseFail(g.w, http.StatusNotFound, protocol.ErrorCodeFileNotFound, "默认头像文件不存在")
			return
		}
		http.ServeFile(g.w, g.r, avatarPath)
		return
	}

	userID, err := g.extractor.GetUserId()
	if err != nil {
		protocol.HttpResponseFail(g.w, http.StatusInternalServerError, protocol.ErrorCodeMissingUserId, fmt.Sprintf("%v", err))
		return
	}

	realAvatarPath := filepath.Join(avatarDirPath, fmt.Sprintf("%v", userID), avatarFileName)
	if _, err := os.Stat(realAvatarPath); os.IsNotExist(err) {
		protocol.HttpResponseFail(g.w, http.StatusNotFound, protocol.ErrorCodeFileNotFound, "用户头像文件不存在")
		return
	}
	http.ServeFile(g.w, g.r, realAvatarPath)
}
