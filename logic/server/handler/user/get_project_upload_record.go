package user

import (
	"code-comment-analyzer/protocol"
	"fmt"
	"net/http"
	"strconv"

	"code-comment-analyzer/data"
	"code-comment-analyzer/server/middleware"
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
	// 从 URL 查询参数中获取分页参数
	pageStr := g.r.URL.Query().Get("page")
	perPageStr := g.r.URL.Query().Get("perPage")

	// 将分页参数从字符串转换为整数
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		protocol.HttpResponseFail(g.w, http.StatusBadRequest, protocol.ErrorCodeBadRequest, "无效的页码")
		return
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		protocol.HttpResponseFail(g.w, http.StatusBadRequest, protocol.ErrorCodeBadRequest, "无效的每页数量")
		return
	}

	// 获取操作管理器实例
	om := g.registry.GetOperationManager()

	// 根据分页参数获取用户操作记录列表以及总记录数
	records, total, err := om.GetUserOperatingRecords(page, perPage)
	if err != nil {
		protocol.HttpResponseFail(g.w, http.StatusInternalServerError, protocol.ErrorCodeInternalServerError, fmt.Sprintf("%v", err))
		return
	}

	// 构造返回数据
	responseData := map[string]interface{}{
		"data":  records,
		"total": total,
		"page":  page,
	}

	protocol.HttpResponseSuccess(g.w, http.StatusOK, "获取上传操作记录成功", protocol.WithData(responseData))
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
