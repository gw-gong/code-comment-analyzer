package main

import (
	"code-comment-analyzer/ccanalyzer_client"
	"code-comment-analyzer/data/redis"
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
	defer mysqlMaster.Close()
	exitOnErr(err)
	register.Register(mysqlMaster)

	redisMaster := redis.NewSessionManager(cfg.RedisMaster, cfg.UserTokenDuration)
	defer redisMaster.Close()
	register.Register(redisMaster)

	ccanalyzer, err := ccanalyzer_client.NewCCAnalyzer(cfg.CcAnalyzerConfig)
	defer ccanalyzer.Close()
	exitOnErr(err)

	httpServer := server.NewHTTPServer()
	httpServer.RegisterRouters(register, ccanalyzer)
	httpServerCfg := config.Cfg.HttpServerConfig
	httpServer.Listen(httpServerCfg.Host, httpServerCfg.Port)
}

func exitOnErr(err error) {
	if err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
}
