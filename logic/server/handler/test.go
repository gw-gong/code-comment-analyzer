package handler

import (
	"code-comment-analyzer/ccanalyzer_client"
	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
	"fmt"
	"log"
	"net/http"
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
		protocol.HttpResponseFail(t.w, http.StatusInternalServerError, protocol.ErrorCodeMissingUserId, fmt.Sprintf("%v", err))
		return
	}
	log.Printf("TestXXX.handle()|%d", userID)

	var (
		sqlExecutor = t.registry.GetTestSqlExecutor()
		ccanalyzer  = t.ccanalyzer
	)

	// test SQL
	err = sqlExecutor.InsertXXX()
	if err != nil {
		protocol.HttpResponseFail(t.w, http.StatusInternalServerError, protocol.ErrorCodeRPCCallFail, fmt.Sprintf("%v", err))
		return
	}
	log.Printf("Insertxxx Successfully")

	// test RPC call
	analyzedData, err := ccanalyzer.AnalyzeFileContent("Python", "# 这是一个注释\n")
	if err != nil {
		protocol.HttpResponseFail(t.w, http.StatusInternalServerError, protocol.ErrorCodeRPCCallFail, fmt.Sprintf("%v", err))
		return
	}

	// 设置HTTP头部的Content-Type为text/plain，表示发送的是纯文本
	t.w.Header().Set("Content-Type", "application/json")

	// 写入响应状态码（HTTP 200 OK）
	t.w.WriteHeader(http.StatusOK)
	// 向响应体写入一条消息
	_, err = fmt.Fprintln(t.w, "This is a test route handler function. Insert successfully\n"+analyzedData)
	if err != nil {
		log.Printf("Error writing response:%v", err)
	}
}
