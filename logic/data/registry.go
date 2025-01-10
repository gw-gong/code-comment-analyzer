package data

import (
	"code-comment-analyzer/data/mysql"
	"fmt"
)

var (
	ErrUnknownDataManager  = fmt.Errorf("unknown Data Manager")
	ErrDataManagerNotFound = fmt.Errorf("data Manager Not Found")
)

type DataManagerRegistry struct {
	sqlExecutor mysql.SqlExecutor
}

func (registry *DataManagerRegistry) Register(elem interface{}) {
	switch elem.(type) {
	case mysql.SqlExecutor:
		registry.sqlExecutor = elem.(mysql.SqlExecutor)
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
