package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"code-comment-analyzer/data/mysql/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type mysqlClient struct {
	db *sql.DB
}

func initMysqlClient(host, port, userName, password, dbName string) (*mysqlClient, error) {
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

func (master *mysqlClient) GetUserInfoByEmail(email string) (userID uint64, nickname string, password string, err error) {
	var queryMods []qm.QueryMod
	queryMods = append(queryMods, qm.Select(models.UserUserColumns.UID, models.UserUserColumns.Nickname, models.UserUserColumns.Password))
	queryMods = append(queryMods, models.UserUserWhere.Email.EQ(email))
	user, err := models.UserUsers(queryMods...).One(master.db)
	if err != nil {
		return
	}
	if user == nil {
		err = fmt.Errorf("user not found")
		return
	}
	return user.UID, user.Nickname, user.Password, err
}
