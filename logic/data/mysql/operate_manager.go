package mysql

import (
	"database/sql"

	"code-comment-analyzer/config"
	"code-comment-analyzer/protocol"
)

type OperationManager interface {
	createOperate(tx *sql.Tx, userID uint64, operationType string) (operateID int64, err error)

	RecordFileUpload(userID uint64, language, fileContent string) (err error)
	RecordProjectUpload(userID uint64, projectUrl string) (err error)

	GetOneProjectUploadRecordUrlByOpID(operatingRecordId int64) (projectUrl string, err error)
	GetUserOperatingRecords(page, perPage int) (records []protocol.OperatingRecord, total int64, err error)
	GetOneFileUploadRecordByOpID(operatingRecordId int64) (language string, fileContent string, err error)

	DeleteOperatingRecordByID(operatingRecordId int64) (err error)
	
	Close()
}

func NewOperationManager(cfgMaster config.MysqlConfig) (OperationManager, error) {
	return initMysqlClient(cfgMaster.Host, cfgMaster.Port, cfgMaster.Username, cfgMaster.Password, cfgMaster.DBName)
}

const (
	OperationTypeFileUpload    = "文件上传"
	OperationTypeProjectUpload = "项目上传"
)
