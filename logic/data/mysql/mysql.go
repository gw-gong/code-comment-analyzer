package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"code-comment-analyzer/data/mysql/models"
	"code-comment-analyzer/protocol"

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
func (m *mysqlClient) CreateUser(email string, password string, nickname string) (uint64, error) {
	// 创建一个新的用户对象
	user := models.UserUser{
		Email:      email,
		Password:   password, // 明文存储密码
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
	return user.UID, nil
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

func (m *mysqlClient) GetUserInfoByUserID(userID uint64) (isSetProfilePicture bool, profilePictureUrl string, email string, nickname string, dateJoined time.Time, err error) {
	var queryMods []qm.QueryMod
	queryMods = append(queryMods, qm.Select(models.UserUserColumns.ProfilePicture, models.UserUserColumns.Email, models.UserUserColumns.Nickname, models.UserUserColumns.DateJoined))
	queryMods = append(queryMods, models.UserUserWhere.UID.EQ(userID))
	user, err := models.UserUsers(queryMods...).One(m.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("user not found|userID:%v", userID)
			return
		}
		return
	}
	if user.ProfilePicture.Valid == false {
		return false, "", user.Email, user.Nickname, user.DateJoined, nil
	}
	return true, user.ProfilePicture.String, user.Email, user.Nickname, user.DateJoined, nil
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

func (m *mysqlClient) UpdatePassword(userID uint64, oldPassword, newPassword string) error {
	// 查找用户
	var queryMods []qm.QueryMod
	queryMods = append(queryMods, models.UserUserWhere.UID.EQ(userID))
	user, err := models.UserUsers(queryMods...).One(m.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 用户不存在
			return errors.New("用户不存在")
		}
		// 其他错误
		return err
	}

	// 验证旧密码是否正确
	if user.Password != oldPassword {
		// 旧密码不匹配
		return errors.New("旧密码错误")
	}

	// 更新密码
	user.Password = newPassword
	_, err = user.Update(m.db, boil.Infer())
	if err != nil {
		// 更新密码失败
		log.Printf("Error updating password: %v\n", err)
		return errors.New("密码更新失败")
	}

	// 密码更新成功
	return nil
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

func (m *mysqlClient) GetOneProjectUploadRecordUrlByOpID(operatingRecordId int64) (projectUrl string, err error) {
	var queryMods []qm.QueryMod
	queryMods = append(queryMods, qm.Select(models.UserProjectrecordColumns.ProjectURL))
	queryMods = append(queryMods, models.UserProjectrecordWhere.OperatingRecordID.EQ(operatingRecordId))
	projectRecord, err := models.UserProjectrecords(queryMods...).One(m.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = fmt.Errorf("project record not found|operatingRecordId:%v", operatingRecordId)
			return
		}
		return
	}
	return projectRecord.ProjectURL, nil
}

func (m *mysqlClient) DeleteOperatingRecordByID(operatingRecordId int64) (err error) {
	// 开启事务
	tx, err := m.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// 首先获取操作记录
	opRecord, err := models.UserOperatingrecords(
		models.UserOperatingrecordWhere.ID.EQ(operatingRecordId),
	).One(tx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("operation record not found|recordID:%v", operatingRecordId)
		}
		return err
	}

	// 根据操作类型删除对应的记录
	switch opRecord.OperationType {
	case OperationTypeFileUpload:
		_, err = models.UserFilerecords(
			models.UserFilerecordWhere.OperatingRecordID.EQ(operatingRecordId),
		).DeleteAll(tx)
	case OperationTypeProjectUpload:
		_, err = models.UserProjectrecords(
			models.UserProjectrecordWhere.OperatingRecordID.EQ(operatingRecordId),
		).DeleteAll(tx)
	default:
		return fmt.Errorf("unknown operation type|operationType:%v", opRecord.OperationType)
	}
	if err != nil {
		return fmt.Errorf("failed to delete child record: %v", err)
	}

	// 最后删除操作记录
	_, err = opRecord.Delete(tx)
	if err != nil {
		return fmt.Errorf("failed to delete operation record|recordID:%v, err:%v", operatingRecordId, err)
	}

	return nil
}

func (m *mysqlClient) GetUserOperatingRecords(page, perPage int) (records []protocol.OperatingRecord, total int64, err error) {
	// 设置分页查询参数
	offset := (page - 1) * perPage
	queryMods := []qm.QueryMod{
		qm.Limit(perPage),
		qm.Offset(offset),
	}

	// 执行查询获取操作记录列表
	operatingRecords, err := models.UserOperatingrecords(queryMods...).All(m.db)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to fetch user operating records: %v", err)
	}

	// 构建返回的操作记录列表
	records = make([]protocol.OperatingRecord, 0, len(operatingRecords))
	for _, record := range operatingRecords {
		records = append(records, protocol.OperatingRecord{
			ID:            record.ID,
			OperationType: record.OperationType,
			CreatedAt:     record.CreatedAt.Format("2006-01-02"),
			UpdatedAt:     record.UpdatedAt.Format("2006-01-02"),
		})
	}

	return records, int64(len(records)), nil
}

func (om *mysqlClient) GetOneFileUploadRecordByOpID(operatingRecordId int64) (language string, fileContent string, err error) {
	var queryMods []qm.QueryMod
	queryMods = append(queryMods, qm.Select(models.UserFilerecordColumns.FileType, models.UserFilerecordColumns.FileContent))
	queryMods = append(queryMods, models.UserFilerecordWhere.OperatingRecordID.EQ(operatingRecordId))

	fileRecord, err := models.UserFilerecords(queryMods...).One(om.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", fmt.Errorf("file record not found|operatingRecordId:%v", operatingRecordId)
		}
		return "", "", err
	}

	return fileRecord.FileType, fileRecord.FileContent, nil
}

func (m *mysqlClient) UpdateUserAvatar(userID uint64, avatarFileName string) error {
	if avatarFileName == "" {
		fmt.Println("avatarFileName is empty")
		return nil
	}

	var queryMods []qm.QueryMod
	queryMods = append(queryMods, models.UserUserWhere.UID.EQ(userID))

	user, err := models.UserUsers(queryMods...).One(m.db)
	if err != nil {
		return fmt.Errorf("failed to find user: %v", err)
	}

	user.ProfilePicture.String = avatarFileName
	user.ProfilePicture.Valid = true
	_, err = user.Update(m.db, boil.Infer())
	return err
}

func (m *mysqlClient) UpdateUserInfo(userID uint64, nickname string) error {
	var queryMods []qm.QueryMod
	queryMods = append(queryMods, models.UserUserWhere.UID.EQ(userID))

	user, err := models.UserUsers(queryMods...).One(m.db)
	if err != nil {
		return fmt.Errorf("failed to find user: %v", err)
	}

	user.Nickname = nickname

	_, err = user.Update(m.db, boil.Infer())
	return err
}
