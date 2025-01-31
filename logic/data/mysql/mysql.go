package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"code-comment-analyzer/config"
	"code-comment-analyzer/data/mysql/models"
)

type TestSqlExecutor interface {
	InsertXXX() error
	Close()
}

func NewTestSqlExecutor(cfgMaster config.MysqlConfig) (TestSqlExecutor, error) {
	return initMysqlMaster(cfgMaster.Host, cfgMaster.Port, cfgMaster.Username, cfgMaster.Password, cfgMaster.DBName)
}

type mysqlClient struct {
	db *sql.DB
}

func initMysqlMaster(host, port, userName, password, dbName string) (*mysqlClient, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", userName, password, host, port, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	log.Println(dsn)

	return &mysqlClient{db: db}, nil
}

func (master *mysqlClient) Close() {
	err := master.db.Close()
	if err != nil {
		panic(err)
	}
}

func (master *mysqlClient) InsertXXX() error {
	user := models.UserUser{
		Email:      fmt.Sprintf("xpl%d@ccanalyzer.com", rand.Int()),
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
