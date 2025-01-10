package handler

import (
	"fmt"
	"log"
	"net/http"

	"code-comment-analyzer/ccanalyzer_client"
	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
)

type TestXXX struct {
	w          http.ResponseWriter
	r          *http.Request
	extractor  middleware.Extractor
	registry   *data.DataManagerRegistry
	ccanalyzer ccanalyzer_client.CCAnalyzer
}

func NewTestXXX(registry *data.DataManagerRegistry, ccanalyzer ccanalyzer_client.CCAnalyzer) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &TestXXX{
			w:          w,
			r:          r,
			extractor:  extractor,
			registry:   registry,
			ccanalyzer: ccanalyzer,
		}
	}
}

func (t *TestXXX) Handle() {
	userID, err := t.extractor.GetUserId()
	if err != nil {
		protocol.HandleError(t.w, protocol.ErrorCodeMissingUserId, err)
		return
	}
	log.Printf("TestXXX.handle()|%d", userID)

	var (
		sqlExecutor = t.registry.GetSqlExecutor()
		ccanalyzer  = t.ccanalyzer
	)

	// test SQL
	err = sqlExecutor.InsertXXX()
	if err != nil {
		return
	}
	log.Printf("Insertxxx Successfully")

	// test RPC call
	resp, err := ccanalyzer.AddUser("xpl", 23)
	if err != nil {
		protocol.HandleError(t.w, protocol.ErrorCodeRPCCallFail, err)
		return
	}
	log.Printf("rpc call successfully | call back: %v", resp)

	// 设置HTTP头部的Content-Type为text/plain，表示发送的是纯文本
	t.w.Header().Set("Content-Type", "text/plain")
	// 写入响应状态码（HTTP 200 OK）
	t.w.WriteHeader(http.StatusOK)
	// 向响应体写入一条消息
	_, err = fmt.Fprintln(t.w, "This is a test route handler function. Insert successfully"+resp.GetMsg(), resp.GetCode())
	if err != nil {
		log.Printf("Error writing response:", err)
	}
}
