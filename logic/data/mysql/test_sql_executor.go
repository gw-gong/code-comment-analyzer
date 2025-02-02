package mysql

import "code-comment-analyzer/config"

type TestSqlExecutor interface {
	InsertXXX() error
	Close()
}

func NewTestSqlExecutor(cfgMaster config.MysqlConfig) (TestSqlExecutor, error) {
	return initMysqlClient(cfgMaster.Host, cfgMaster.Port, cfgMaster.Username, cfgMaster.Password, cfgMaster.DBName)
}
