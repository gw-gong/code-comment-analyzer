package file_storage

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
	"code-comment-analyzer/util"
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
	avatarFileName := filepath.Base(g.r.URL.Path)
	userID, err := g.extractor.GetUserId()
	if err != nil {
		protocol.HttpResponseFail(g.w, http.StatusInternalServerError, protocol.ErrorCodeMissingUserId, fmt.Sprintf("%v", err))
		return
	}

	realAvatarPath := util.GetAvatarStoragePath(userID, avatarFileName)
	if _, err = os.Stat(realAvatarPath); os.IsNotExist(err) {
		err = fmt.Errorf("用户头像文件不存在")
		log.Printf("Error|GetAvatars|realAvatarPath:%s|err: %v", realAvatarPath, err)
		protocol.HttpResponseFail(g.w, http.StatusNotFound, protocol.ErrorCodeFileNotFound, fmt.Sprintf("%v", err))
		return
	}
	http.ServeFile(g.w, g.r, realAvatarPath)
}
