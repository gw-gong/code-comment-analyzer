package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"code-comment-analyzer/config"
	"code-comment-analyzer/data/mysql/models"
)

type Executor interface {
	InsertXXX() error
}

func GetMysqlMasterExecutor(cfgMaster config.MysqlConfig) (Executor, error) {
	return initMysqlMaster(cfgMaster.Host, cfgMaster.Port, cfgMaster.Username, cfgMaster.Password, cfgMaster.DBName)
}

type MysqlMaster struct {
	db *sql.DB
}

func initMysqlMaster(host, port, userName, password, dbName string) (*MysqlMaster, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", userName, password, host, port, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	fmt.Println(dsn)

	return &MysqlMaster{db: db}, nil
}

func (master *MysqlMaster) CloseMysqlMaster() {
	err := master.db.Close()
	if err != nil {
		panic(err)
	}
}

func (master *MysqlMaster) InsertXXX() error {
	user := models.User{
		Username: "username1",
		Password: "123456",
		NickName: "nickName1",
	}
	err := user.Insert(master.db, boil.Infer())
	if err != nil {
		fmt.Printf("Error inserting user: %v\n", err)
		return err
	}

	fmt.Printf("User inserted: %+v\n", user)
	return nil
}
