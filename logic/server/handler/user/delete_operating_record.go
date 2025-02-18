package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
)

type DeleteOperatingRecord struct {
	w         http.ResponseWriter
	r         *http.Request
	extractor middleware.Extractor
	registry  *data.DataManagerRegistry
}

func NewDeleteOperatingRecord(registry *data.DataManagerRegistry) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &DeleteOperatingRecord{
			w:         w,
			r:         r,
			extractor: extractor,
			registry:  registry,
		}
	}
}

func (d *DeleteOperatingRecord) Handle() {
	requestData, err := d.decodeRequest()
	if err != nil {
		return
	}

	if requestData.ID == 0 {
		protocol.HttpResponseFail(d.w, http.StatusBadRequest, protocol.ErrorCodeInvalidID, "Invalid operation ID")
		return
	}

	om := d.registry.GetOperationManager()
	// Delete the operating record by ID
	err = om.DeleteOperatingRecordByID(requestData.ID)
	if err != nil {
		protocol.HttpResponseFail(d.w, http.StatusInternalServerError, protocol.ErrorCodeInternalServerError, fmt.Sprintf("Failed to delete operating record: %v", err))
		return
	}

	protocol.HttpResponseSuccess(d.w, http.StatusOK, "删除操作记录成功")
}

func (d *DeleteOperatingRecord) decodeRequest() (*protocol.DeleteOperatingRecordRequest, error) {
	var requestData protocol.DeleteOperatingRecordRequest
	// 解码请求体的 JSON 数据到 requestData 结构体
	err := json.NewDecoder(d.r.Body).Decode(&requestData)
	if err != nil {
		log.Println("Failed to parse JSON body:", err)
		protocol.HttpResponseFail(d.w, http.StatusBadRequest, protocol.ErrorCodeParseRequestFailed, "Invalid JSON format")
		return nil, err
	}
	return &requestData, nil
}
