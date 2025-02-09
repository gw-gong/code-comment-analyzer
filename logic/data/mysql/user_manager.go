package mysql

import (
	"code-comment-analyzer/config"
	"time"
)

type UserManager interface {
	CreateUser(email string, password string, nickname string) (uint64, error)

	GetUserInfoByEmail(email string) (userID uint64, nickname string, password string, err error)
	GetUserInfoByUserID(userID uint64) (email string, nickname string, dateJoined time.Time, err error)
	GetUserProfilePictureByUserID(userID uint64) (isSetProfilePicture bool, profilePictureUrl string, err error)
	IsExistUserByEmail(email string) (isExist bool, err error)

	Close()
}

func NewUserManager(cfgMaster config.MysqlConfig) (UserManager, error) {
	return initMysqlClient(cfgMaster.Host, cfgMaster.Port, cfgMaster.Username, cfgMaster.Password, cfgMaster.DBName)
}
