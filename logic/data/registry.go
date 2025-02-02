package data

import (
	"code-comment-analyzer/data/mysql"
	"code-comment-analyzer/data/redis"
	"fmt"
	"log"
)

var (
	ErrUnknownDataManager  = fmt.Errorf("unknown Data Manager")
	ErrDataManagerNotFound = fmt.Errorf("data Manager Not Found")
)

type DataManagerRegistry struct {
	testSqlExecutor  mysql.TestSqlExecutor
	userManager      mysql.UserManager
	operationManager mysql.OperationManager
	sessionManager   redis.SessionManager
}

func (registry *DataManagerRegistry) RegisterTestSqlExecutor(testSqlExecutor mysql.TestSqlExecutor) {
	registry.testSqlExecutor = testSqlExecutor
	log.Println("Registered TestSqlExecutor")
}

func (registry *DataManagerRegistry) GetTestSqlExecutor() mysql.TestSqlExecutor {
	if registry.testSqlExecutor != nil {
		return registry.testSqlExecutor
	}
	panic(ErrDataManagerNotFound)
}

func (registry *DataManagerRegistry) RegisterUserManager(userManager mysql.UserManager) {
	registry.userManager = userManager
	log.Println("Registered UserManager")
}

func (registry *DataManagerRegistry) GetUserManager() mysql.UserManager {
	if registry.userManager != nil {
		return registry.userManager
	}
	panic(ErrDataManagerNotFound)
}

func (registry *DataManagerRegistry) RegisterOperationManager(operationManager mysql.OperationManager) {
	registry.operationManager = operationManager
	log.Println("Registered OperationManager")
}

func (registry *DataManagerRegistry) GetOperationManager() mysql.OperationManager {
	if registry.operationManager != nil {
		return registry.operationManager
	}
	panic(ErrDataManagerNotFound)
}

func (registry *DataManagerRegistry) RegisterSessionManager(sessionManager redis.SessionManager) {
	registry.sessionManager = sessionManager
	log.Println("Registered SessionManager")
}

func (registry *DataManagerRegistry) GetSessionManager() redis.SessionManager {
	if registry.sessionManager != nil {
		return registry.sessionManager
	}
	panic(ErrDataManagerNotFound)
}
