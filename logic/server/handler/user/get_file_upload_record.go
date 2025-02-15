package user

import (
	"fmt"
	"net/http"

	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
)

type GetFileUploadRecord struct {
	w         http.ResponseWriter
	r         *http.Request
	extractor middleware.Extractor
	registry  *data.DataManagerRegistry
}

func NewGetFileUploadRecord(registry *data.DataManagerRegistry) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &GetFileUploadRecord{
			w:         w,
			r:         r,
			extractor: extractor,
			registry:  registry,
		}
	}
}

func (g *GetFileUploadRecord) Handle() {
	operatingRecordId, err := g.decodeRequest()
	if err != nil {
		return
	}

	om := g.registry.GetOperationManager()
	language, fileContent, err := om.GetOneFileUploadRecordByOpID(operatingRecordId)
	if err != nil {
		protocol.HttpResponseFail(g.w, http.StatusInternalServerError, protocol.ErrorCodeInternalServerError, fmt.Sprintf("%v", err))
		return
	}

	response := &protocol.GetFileUploadRecordResponse{
		Language:    language,
		FileContent: fileContent,
	}

	protocol.HttpResponseSuccess(g.w, http.StatusOK, "获取文件上传记录成功", protocol.WithData(response))
}

func (g *GetFileUploadRecord) decodeRequest() (operatingRecordId int64, err error) {
	id := g.r.URL.Query().Get(protocol.GetKeyOperatingRecordId)
	if id == "" {
		err = fmt.Errorf("operatingRecordId is missing")
		protocol.HttpResponseFail(g.w, http.StatusBadRequest, protocol.ErrorCodeParamError, fmt.Sprintf("%v", err))
		return
	}
	return protocol.OpIDTransformStr2Int64(id), nil
}
