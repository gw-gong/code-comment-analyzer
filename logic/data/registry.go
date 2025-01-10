package data

import (
	"code-comment-analyzer/data/redis"
	"fmt"

	"code-comment-analyzer/data/mysql"
)

var (
	ErrUnknownDataManager  = fmt.Errorf("unknown Data Manager")
	ErrDataManagerNotFound = fmt.Errorf("data Manager Not Found")
)

type DataManagerRegistry struct {
	sqlExecutor    mysql.SqlExecutor
	sessionManager redis.SessionManager
}

func (registry *DataManagerRegistry) Register(elem interface{}) {
	switch elem.(type) {
	case mysql.SqlExecutor:
		registry.sqlExecutor = elem.(mysql.SqlExecutor)
	case redis.SessionManager:
		registry.sessionManager = elem.(redis.SessionManager)
	default:
		panic(ErrUnknownDataManager)
	}
}

func (registry *DataManagerRegistry) GetSqlExecutor() mysql.SqlExecutor {
	if registry.sqlExecutor != nil {
		return registry.sqlExecutor
	}
	panic(ErrDataManagerNotFound)
}

func (registry *DataManagerRegistry) GetSessionManager() redis.SessionManager {
	if registry.sessionManager != nil {
		return registry.sessionManager
	}
	panic(ErrDataManagerNotFound)
}
