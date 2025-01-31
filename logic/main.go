package main

import (
	"log"
	"os"

	"code-comment-analyzer/ccanalyzer_client"
	"code-comment-analyzer/config"
	"code-comment-analyzer/data"
	"code-comment-analyzer/data/mysql"
	"code-comment-analyzer/data/redis"
	"code-comment-analyzer/server"
)

func main() {
	cfg := config.Cfg

	register := &data.DataManagerRegistry{}

	testSqlExecutor, err := mysql.NewTestSqlExecutor(cfg.MysqlMaster)
	defer testSqlExecutor.Close()
	exitOnErr(err)
	register.Register(testSqlExecutor)

	sessionManager := redis.NewSessionManager(cfg.RedisMaster, cfg.UserTokenDuration)
	defer sessionManager.Close()
	register.Register(sessionManager)

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
		log.Printf("%+v", err.Error())
		os.Exit(1)
	}
}
