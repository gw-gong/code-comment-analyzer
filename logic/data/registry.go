package data

import (
	"fmt"

	"code-comment-analyzer/data/mysql"
	"code-comment-analyzer/data/redis"
)

var (
	ErrUnknownDataManager  = fmt.Errorf("unknown Data Manager")
	ErrDataManagerNotFound = fmt.Errorf("data Manager Not Found")
)

type DataManagerRegistry struct {
	testSqlExecutor mysql.TestSqlExecutor
	userManager     mysql.UserManager
	sessionManager  redis.SessionManager
}

func (registry *DataManagerRegistry) Register(elem interface{}) {
	switch elem.(type) {
	case mysql.TestSqlExecutor:
		registry.testSqlExecutor = elem.(mysql.TestSqlExecutor)
	case mysql.UserManager:
		registry.userManager = elem.(mysql.UserManager)
	case redis.SessionManager:
		registry.sessionManager = elem.(redis.SessionManager)
	default:
		panic(ErrUnknownDataManager)
	}
}

func (registry *DataManagerRegistry) GetTestSqlExecutor() mysql.TestSqlExecutor {
	if registry.testSqlExecutor != nil {
		return registry.testSqlExecutor
	}
	panic(ErrDataManagerNotFound)
}

func (registry *DataManagerRegistry) GetUserManager() mysql.UserManager {
	if registry.userManager != nil {
		return registry.userManager
	}
	panic(ErrDataManagerNotFound)
}

func (registry *DataManagerRegistry) GetSessionManager() redis.SessionManager {
	if registry.sessionManager != nil {
		return registry.sessionManager
	}
	panic(ErrDataManagerNotFound)
}
