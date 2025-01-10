package handler

import (
	"code-comment-analyzer/data"
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/middleware"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"

	"code-comment-analyzer/ccanalyzer_client"
)

type TestXXX struct {
	w         http.ResponseWriter
	r         *http.Request
	extractor middleware.Extractor
	registry  *data.DataManagerRegistry
}

func NewTestXXX(registry *data.DataManagerRegistry) middleware.GetHandler {
	return func(w http.ResponseWriter, r *http.Request, extractor middleware.Extractor) middleware.Handler {
		return &TestXXX{
			w:         w,
			r:         r,
			extractor: extractor,
			registry:  registry,
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

	// 连接到server端，此处禁用安全传输，没有加密和验证
	conn, err := grpc.Dial("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// 建立连接
	client := ccanalyzer_client.NewCcAnalyzerClient(conn)
	// 执行rpc调用（这个方法在服务端来实现并返回结果）
	resp, err := client.AddUser(context.Background(), &ccanalyzer_client.UserRequest{Name: "xpl", Age: 23})
	if err != nil {
		protocol.HandleError(t.w, protocol.ErrorCodeRPCCallFail, err)
		return
	}

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
