package server_info

import (
	"os"
)

var (
	serverRunningPath string
)

func init() {
	var err error
	serverRunningPath, err = os.Getwd()
	if err != nil {
		panic("获取服务器运行路径失败")
		return
	}
}

func GetServerRunningPath() string {
	return serverRunningPath
}
