package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"code-comment-analyzer/config"
	"code-comment-analyzer/data/mysql/models"
)

type SqlExecutor interface {
	InsertXXX() error
	Close()
}

func GetMysqlMasterExecutor(cfgMaster config.MysqlConfig) (SqlExecutor, error) {
	return initMysqlMaster(cfgMaster.Host, cfgMaster.Port, cfgMaster.Username, cfgMaster.Password, cfgMaster.DBName)
}

type mysqlMaster struct {
	db *sql.DB
}

func initMysqlMaster(host, port, userName, password, dbName string) (*mysqlMaster, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", userName, password, host, port, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	fmt.Println(dsn)

	return &mysqlMaster{db: db}, nil
}

func (master *mysqlMaster) Close() {
	err := master.db.Close()
	if err != nil {
		panic(err)
	}
}

func (master *mysqlMaster) InsertXXX() error {
	user := models.UserUser{
		Email:      "xpl111@ccanalyzer.com",
		Password:   "123456",
		Nickname:   "xpl",
		DateJoined: time.Now(),
		IsActive:   true,
	}
	err := user.Insert(master.db, boil.Infer())
	if err != nil {
		fmt.Printf("Error inserting user: %v\n", err)
		return err
	}

	fmt.Printf("User inserted: %+v\n", user)
	return nil
}
