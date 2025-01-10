package main

import (
	"log"
	"os"

	"code-comment-analyzer/config"
	"code-comment-analyzer/data"
	"code-comment-analyzer/data/mysql"
	"code-comment-analyzer/server"
)

func main() {
	cfg := config.Cfg

	register := &data.DataManagerRegistry{}
	mysqlMaster, err := mysql.GetMysqlMasterExecutor(cfg.MysqlMaster)
	exitOnErr(err)
	register.Register(mysqlMaster)

	httpServer := server.NewHTTPServer()
	httpServer.RegisterRouters(register)
	httpServerCfg := config.Cfg.HttpServerConfig
	httpServer.Listen(httpServerCfg.Host, httpServerCfg.Port)
}

func exitOnErr(err error) {
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
}
