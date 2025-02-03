package mysql

import (
	"code-comment-analyzer/config"
	"code-comment-analyzer/data/mysql/models"
)

type UserManager interface {
	GetUserInfoByEmail(email string) (userID uint64, nickname string, password string, err error)
	GetUserByEmail(email string) (*models.UserUser, error)
	CreateUser(email string, password string, nickname string) (uint64, error)
	Close()
}

func NewUserManager(cfgMaster config.MysqlConfig) (UserManager, error) {
	return initMysqlClient(cfgMaster.Host, cfgMaster.Port, cfgMaster.Username, cfgMaster.Password, cfgMaster.DBName)
}
