package main

import (
	"code-comment-analyzer/config"
	"code-comment-analyzer/server"
)

func main() {
	httpServer := server.NewHTTPServer()
	httpServer.RegisterRouters()
	httpServerCfg := config.Cfg.HttpServerConfig
	httpServer.Listen(httpServerCfg.Host, httpServerCfg.Port)
}
