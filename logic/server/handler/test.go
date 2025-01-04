package handler

import (
	"fmt"
	"net/http"

	"code-comment-analyzer/config"
	"code-comment-analyzer/data/mysql"
)

// Test 是一个HTTP处理函数，它接收两个参数：http.ResponseWriter 和 *http.Request
func Test(w http.ResponseWriter, r *http.Request) {
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

	// 设置HTTP头部的Content-Type为text/plain，表示发送的是纯文本
	w.Header().Set("Content-Type", "text/plain")

	// 写入响应状态码（HTTP 200 OK）
	w.WriteHeader(http.StatusOK)

	// 向响应体写入一条消息
	_, err = fmt.Fprintln(w, "This is a test route handler function. Insert successfully")
	if err != nil {
		// 如果写入响应时发生错误，可以在服务器日志中记录此错误
		// 这里简单打印到标准错误输出，你也可以使用更复杂的日志记录方式
		fmt.Println("Error writing response:", err)
	}
}
