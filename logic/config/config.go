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

type RedisConfig struct {
	Host          string `yaml:"host"`
	Port          string `yaml:"port"`
	Password      string `yaml:"password"`
	DBNum         int    `yaml:"db"`
	PrefixSession string `yaml:"prefix_session"`
}

type HttpServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type CcAnalyzerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type FileStoragePathConfig struct {
	Avatars  string `yaml:"avatars"`
	Projects string `yaml:"projects"`
}

type Config struct {
	HttpServerConfig  HttpServerConfig      `yaml:"http_server"`
	MysqlMaster       MysqlConfig           `yaml:"mysql_master"`
	RedisMaster       RedisConfig           `yaml:"redis_master"`
	CcAnalyzerConfig  CcAnalyzerConfig      `yaml:"ccanalyzer_conf"`
	UserTokenDuration uint32                `yaml:"user_token_duration"`
	JwtKey            string                `yaml:"jwt_key"`
	DefaultNickname   string                `yaml:"default_nickname"`
	DefaultAvatar     string                `yaml:"default_avatar"`
	FileStoragePath   FileStoragePathConfig `yaml:"file_storage_path"`
	ReadmePath        string                `yaml:"code-comment-analyzer_readme_path"`
	MaxFileSize       int64                 `yaml:"max_file_size"`
	MaxProjectSize    int64                 `yaml:"max_project_size"`
	MaxAvatarSize     int64                 `yaml:"max_avatar_size"`
	UuidProjectPrefix string                `yaml:"uuid_project_prefix"`
	UuidAvatarPrefix  string                `yaml:"uuid_avatar_prefix"`
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

	// print to console
	yamlCfg, err := yaml.Marshal(Cfg)
	if err != nil {
		panic("Error marshalling config to YAML")
	}
	fmt.Println(string(yamlCfg))
}
