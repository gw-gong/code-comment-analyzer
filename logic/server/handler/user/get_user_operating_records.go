package user

import (
	"fmt"
	"net/http"
	"strconv"
	
	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
)

type GetUserOperatingRecords struct {
	w         http.ResponseWriter
	r         *http.Request
	extractor middleware.Extractor
	registry  *data.DataManagerRegistry
}

func NewGetUserOperatingRecords(registry *data.DataManagerRegistry) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &GetUserOperatingRecords{
			w:         w,
			r:         r,
			extractor: extractor,
			registry:  registry,
		}
	}
}

func (g *GetUserOperatingRecords) Handle() {
	page, perPage, err := g.getPageParams()
	if err != nil {
		protocol.HttpResponseFail(g.w, http.StatusBadRequest, protocol.ErrorCodeParamError, fmt.Sprintf("%v", err))
		return
	}

	om := g.registry.GetOperationManager()

	// 根据分页参数获取用户操作记录列表以及总记录数
	records, total, err := om.GetUserOperatingRecords(page, perPage)
	if err != nil {
		protocol.HttpResponseFail(g.w, http.StatusInternalServerError, protocol.ErrorCodeInternalServerError, fmt.Sprintf("%v", err))
		return
	}

	// 构造返回数据
	responseData := protocol.GetUserOperatingRecordsResponse{
		Data:  records,
		Total: total,
		Page:  page,
	}

	protocol.HttpResponseSuccess(g.w, http.StatusOK, "获取上传操作记录成功", protocol.WithData(responseData))
}

func (g *GetUserOperatingRecords) getPageParams() (page int, perPage int, err error) {
	pageStr := g.r.URL.Query().Get(protocol.GetKeyPage)
	perPageStr := g.r.URL.Query().Get(protocol.GetKeyPerPage)

	page, err = strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		err = fmt.Errorf("无效的页码")
		return 0, 0, err
	}

	perPage, err = strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		err = fmt.Errorf("无效的每页数量")
		return 0, 0, err
	}

	return page, perPage, nil
}