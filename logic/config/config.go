package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	Cfg Config
)

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type HttpServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type CcAnalyzerConfig struct {
	Host             string `yaml:"host"`
	Port             string `yaml:"port"`
	ConnectPoolSize  int    `yaml:"connect_pool_size"`
	TimeoutGetClient uint32 `yaml:"timeout_get_client"`
}

type Config struct {
	HttpServerConfig  HttpServerConfig `yaml:"http_server"`
	MysqlMaster       MysqlConfig      `yaml:"mysql_master"`
	CcAnalyzerConfig  CcAnalyzerConfig `yaml:"ccanalyzer_conf"`
	UserTokenDuration uint32           `yaml:"user_token_duration"`
	JwtKey            string           `yaml:"jwt_key"`
}

func init() {
	buf, err := os.ReadFile("config/config.yaml")
	if err != nil {
		panic("read config.yaml fail")
	}

	err = yaml.Unmarshal(buf, &Cfg)
	if err != nil {
		panic("Error parsing config file")
	}

	fmt.Printf("%+v\n", Cfg)
}
