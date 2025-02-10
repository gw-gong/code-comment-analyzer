package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
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

func (m *mysqlClient) Close() {
	err := m.db.Close()
	if err != nil {
		panic(err)
	}
}

func (m *mysqlClient) GetUserInfoByEmail(email string) (userID uint64, nickname string, password string, err error) {
	var queryMods []qm.QueryMod
	queryMods = append(queryMods, qm.Select(models.UserUserColumns.UID, models.UserUserColumns.Nickname, models.UserUserColumns.Password))
	queryMods = append(queryMods, models.UserUserWhere.Email.EQ(email))
	user, err := models.UserUsers(queryMods...).One(m.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("user not found|email:%v", email)
			return
		}
		return
	}
	return user.UID, user.Nickname, user.Password, nil
}

func (m *mysqlClient) GetUserInfoByUserID(userID uint64) (email string, nickname string, dateJoined time.Time, err error) {
	var queryMods []qm.QueryMod
	queryMods = append(queryMods, qm.Select(models.UserUserColumns.Email, models.UserUserColumns.Nickname, models.UserUserColumns.DateJoined))
	queryMods = append(queryMods, models.UserUserWhere.UID.EQ(userID))
	user, err := models.UserUsers(queryMods...).One(m.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("user not found|userID:%v", userID)
			return
		}
		return
	}
	return user.Email, user.Nickname, user.DateJoined, nil
}

func (m *mysqlClient) GetUserProfilePictureByUserID(userID uint64) (isSetProfilePicture bool, profilePictureUrl string, err error) {
	var queryMods []qm.QueryMod
	queryMods = append(queryMods, qm.Select(models.UserUserColumns.ProfilePicture))
	queryMods = append(queryMods, models.UserUserWhere.UID.EQ(userID))
	user, err := models.UserUsers(queryMods...).One(m.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("user not found")
			return
		}
		return
	}
	if user.ProfilePicture.Valid == false {
		return
	}
	return true, user.ProfilePicture.String, nil
}

func (m *mysqlClient) IsExistUserByEmail(email string) (isExist bool, err error) {
	var queryMods []qm.QueryMod
	queryMods = append(queryMods, models.UserUserWhere.Email.EQ(email))
	_, err = models.UserUsers(queryMods...).One(m.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (m *mysqlClient) CreateUser(email string, password string, nickname string) (uint64, error) {
	// 创建一个新的用户对象
	user := models.UserUser{
		Email:      email,
		Password:   password,
		Nickname:   nickname,
		DateJoined: time.Now(),
		IsActive:   true,
	}

	// 将用户数据插入数据库
	err := user.Insert(m.db, boil.Infer())
	if err != nil {
		log.Printf("Error inserting user: %v\n", err)
		return 0, err
	}

	// 返回新创建的用户 ID
	log.Printf("User created: %+v\n", user)
	return user.UID, nil
}

func (m *mysqlClient) createOperate(tx *sql.Tx, userID uint64, operationType string) (operateID int64, err error) {
	op := models.UserOperatingrecord{
		UserID:        userID,
		OperationType: operationType,
	}
	err = op.Insert(tx, boil.Infer()) // 使用事务对象 tx
	if err != nil {
		log.Printf("Error inserting operating record: %v\n", err)
		return -1, err
	}
	return op.ID, nil
}

func (m *mysqlClient) RecordFileUpload(userID uint64, language, fileContent string) (err error) {
	tx, err := m.db.Begin()
	if err != nil {
		log.Printf("RecordFileUpload|Error starting transaction: %v\n", err)
		return err
	}

	// 确保事务最终被提交或回滚
	defer func() {
		if p := recover(); p != nil {
			// 发生 panic，回滚事务
			tx.Rollback()
			panic(p) // 重新抛出 panic
		} else if err != nil {
			// 发生错误，回滚事务
			tx.Rollback()
		} else {
			// 提交事务
			err = tx.Commit()
			if err != nil {
				log.Printf("Error committing transaction: %v\n", err)
			}
		}
	}()

	operationID, err := m.createOperate(tx, userID, OperationTypeFileUpload)
	if err != nil {
		return err
	}

	fileUpload := models.UserFilerecord{
		OperatingRecordID: operationID,
		FileType:          language,
		FileContent:       fileContent,
	}
	err = fileUpload.Insert(tx, boil.Infer())
	if err != nil {
		log.Printf("Error inserting file record: %v\n", err)
		return err
	}
	return nil
}

func (m *mysqlClient) RecordProjectUpload(userID uint64, projectUrl string) (err error) {
	tx, err := m.db.Begin()
	if err != nil {
		log.Printf("RecordProjectUpload|Error starting transaction: %v\n", err)
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
			if err != nil {
				log.Printf("Error committing transaction: %v\n", err)
			}
		}
	}()

	operationID, err := m.createOperate(tx, userID, OperationTypeProjectUpload)
	if err != nil {
		return err
	}

	projectUpload := models.UserProjectrecord{
		OperatingRecordID: operationID,
		ProjectURL:        projectUrl,
	}
	err = projectUpload.Insert(tx, boil.Infer())
	if err != nil {
		log.Printf("Error inserting project record: %v\n", err)
		return err
	}
	return nil
}
