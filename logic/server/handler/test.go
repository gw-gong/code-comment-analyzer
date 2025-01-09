package handler

import (
	"code-comment-analyzer/protocol"
	"code-comment-analyzer/server/jwt"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"

	"code-comment-analyzer/ccanalyzer_client"
	"code-comment-analyzer/config"
	"code-comment-analyzer/data/mysql"
)

// Test 是一个HTTP处理函数，它接收两个参数：http.ResponseWriter 和 *http.Request
func Test(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromRequestCtx(w, r)
	if err != nil {
		protocol.HandleError(w, protocol.ErrorCodeMissingUserId, err)
		return
	}
	fmt.Println(userID)
	exec, err := mysql.GetMysqlMasterExecutor(config.Cfg.MysqlMaster)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	err = exec.InsertXXX()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

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
		protocol.HandleError(w, protocol.ErrorCodeRPCCallFail, err)
		return
	}

	// 设置HTTP头部的Content-Type为text/plain，表示发送的是纯文本
	w.Header().Set("Content-Type", "text/plain")

	// 写入响应状态码（HTTP 200 OK）
	w.WriteHeader(http.StatusOK)

	// 向响应体写入一条消息
	_, err = fmt.Fprintln(w, "This is a test route handler function. Insert successfully"+resp.GetMsg(), resp.GetCode())
	if err != nil {
		// 如果写入响应时发生错误，可以在服务器日志中记录此错误
		// 这里简单打印到标准错误输出，你也可以使用更复杂的日志记录方式
		fmt.Println("Error writing response:", err)
	}
}

func TestLogin(w http.ResponseWriter, r *http.Request) {
	// ....
	jwt.AuthorizeUserToken(123456, w)
}
