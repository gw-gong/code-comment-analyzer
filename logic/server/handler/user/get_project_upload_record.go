package user

import (
	"fmt"
	"net/http"
	"strings"

	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
	"code-comment-analyzer/util"
)

type GetProjectUploadRecord struct {
	w         http.ResponseWriter
	r         *http.Request
	extractor middleware.Extractor
	registry  *data.DataManagerRegistry
}

func NewGetProjectUploadRecord(registry *data.DataManagerRegistry) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &GetProjectUploadRecord{
			w:         w,
			r:         r,
			extractor: extractor,
			registry:  registry,
		}
	}
}

func (g *GetProjectUploadRecord) Handle() {
	operatingRecordId, err := g.decodeRequest()
	if err != nil {
		return
	}

	om := g.registry.GetOperationManager()
	projectUrl, err := om.GetOneProjectUploadRecordUrlByOpID(operatingRecordId)
	if err != nil {
		protocol.HttpResponseFail(g.w, http.StatusInternalServerError, protocol.ErrorCodeInternalServerError, fmt.Sprintf("%v", err))
		return
	}

	directorys := strings.Split(projectUrl, "/")
	if len(directorys) < 2 {
		protocol.HttpResponseFail(g.w, http.StatusInternalServerError, protocol.ErrorCodeInternalServerError, "获取项目名称失败")
		return
	}
	projectStorageName := directorys[len(directorys)-1]

	rootNode := util.BuildDirectoryTree(projectUrl, projectUrl, projectStorageName)

	protocol.HttpResponseSuccess(g.w, http.StatusOK, "获取项目上传记录成功", protocol.WithData(rootNode.Children[0]))
}

func (g *GetProjectUploadRecord) decodeRequest() (operatingRecordId int64, err error) {
	id := g.r.URL.Query().Get(protocol.GetKeyOperatingRecordId)
	if id == "" {
		err = fmt.Errorf("operatingRecordId is missing")
		protocol.HttpResponseFail(g.w, http.StatusBadRequest, protocol.ErrorCodeParamError, fmt.Sprintf("%v", err))
		return
	}
	return protocol.OpIDTransformStr2Int64(id), nil
}