package mysql

import (
	"code-comment-analyzer/config"
	"database/sql"
)

type OperationManager interface {
	createOperate(tx *sql.Tx, userID uint64, operationType string) (operateID int64, err error)

	RecordFileUpload(userID uint64, language, fileContent string) (err error)
	RecordProjectUpload(userID uint64, projectUrl string) (err error)

	GetOneProjectUploadRecordUrlByOpID(operatingRecordId int64) (projectUrl string, err error)

	Close()
}

func NewOperationManager(cfgMaster config.MysqlConfig) (OperationManager, error) {
	return initMysqlClient(cfgMaster.Host, cfgMaster.Port, cfgMaster.Username, cfgMaster.Password, cfgMaster.DBName)
}

const (
	OperationTypeFileUpload    = "文件上传"
	OperationTypeProjectUpload = "项目上传"
)
