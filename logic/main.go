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

	registry := &data.DataManagerRegistry{}

	testSqlExecutor, err := mysql.NewTestSqlExecutor(cfg.MysqlMaster)
	defer testSqlExecutor.Close()
	exitOnErr(err)
	registry.RegisterTestSqlExecutor(testSqlExecutor)

	userManager, err := mysql.NewUserManager(cfg.MysqlMaster)
	defer userManager.Close()
	exitOnErr(err)
	registry.RegisterUserManager(userManager)

	operationManager, err := mysql.NewOperationManager(cfg.MysqlMaster)
	defer operationManager.Close()
	exitOnErr(err)
	registry.RegisterOperationManager(operationManager)

	sessionManager := redis.NewSessionManager(cfg.RedisMaster, cfg.UserTokenDuration)
	defer sessionManager.Close()
	registry.RegisterSessionManager(sessionManager)

	ccanalyzer, err := ccanalyzer_client.NewCCAnalyzer(cfg.CcAnalyzerConfig)
	defer ccanalyzer.Close()
	exitOnErr(err)

	httpServer := server.NewHTTPServer()
	httpServer.RegisterRouters(registry, ccanalyzer)
	httpServerCfg := config.Cfg.HttpServerConfig
	httpServer.Listen(httpServerCfg.Host, httpServerCfg.Port)
}

func exitOnErr(err error) {
	if err != nil {
		log.Printf("%+v", err.Error())
		os.Exit(1)
	}
}
