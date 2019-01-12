package models

import (
	"MPMS/helper"
	"MPMS/structure"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"strings"
)

/**
基础模型 用于写通用方法
@author cyu 2018-4-28 16:29:54
*/
type Model struct {
	Id        int64
	IsDeleted uint8
	CreatorId int64
	CreatedAt string
	UpdatedAt string
}

const (
	//表名定义
	ConfigTableName      = "config"
	UserTableName        = "user"
	MenuTableName        = "menu"
	FlowTableName        = "flow"
	RelationTableName    = "relation"
	CompanyTableName     = "company"
	MiniProgramTableName = "mini_program"

	//删除标志
	UnDeleted = 0
	Deleted   = 1
)

var db *sql.DB
var tx *sql.Tx        //事务使用
var userTrans = false //是否开启事务

/**
连接数据库（使用单例模式）
*/
func (b *Model) initDB() (*sql.DB, error) {
	var err error
	if db != nil {
		return db, err
	}
	fmt.Println("init -- db")
	db, err = sql.Open(beego.AppConfig.String("DBDriverName"), beego.AppConfig.String("DBDataSourceName"))
	if err != nil {
		return db, err
	}
	return db, err
}

/**
开启事务
*/
func (b *Model) StartTrans() (*sql.Tx, error) {
	var err error
	if db == nil {
		db, err = b.initDB()
		if err != nil {
			return nil, err
		}
	}
	if tx, err = db.Begin(); err == nil {
		userTrans = true
	}

	return tx, err
}

/**
事务回滚
*/
func (b *Model) Rollback() error {
	err := tx.Rollback()
	if err == nil {
		userTrans = false
	}
	return err
}

/**
事务回滚
*/
func (b *Model) Commit() error {
	err := tx.Commit()
	if err == nil {
		userTrans = false
	}
	return err
}

/**
查询方法
*/
func (b *Model) query(sql string, args ...interface{}) (*sql.Rows, error) {
	if db == nil {
		if _, err := b.initDB(); err != nil {
			return nil, err
		}
	}

	if userTrans {
		smt, _ := tx.Prepare(sql)
		return smt.Query(args...)
	}

	return db.Query(sql, args...)
}

/**
执行方法
*/
func (b *Model) exec(sql string, args ...interface{}) (sql.Result, error) {

	if db == nil {
		if _, err := b.initDB(); err != nil {
			return nil, err
		}
	}

	if userTrans {
		smt, _ := tx.Prepare(sql)
		return smt.Exec(args...)
	}

	return db.Exec(sql, args...)
}

func (b *Model) quickQuery(fields []string, getFieldsMap func() structure.StringToObjectMap, where structure.StringToObjectMap, table string) (*sql.Rows, structure.Array, error) {
	return b.quickQueryWithExtra(fields, getFieldsMap, where, table, "")
}

func (b *Model) quickQueryWithExtra(fields []string, getFieldsMap func() structure.StringToObjectMap, where structure.StringToObjectMap, table string, extra string) (*sql.Rows, structure.Array, error) {

	//条件字段校验
	var fieldsToCheck []string
	for field := range where {
		fieldsToCheck = append(fieldsToCheck, field)
	}
	if err := b.checkFieldValid(append(fieldsToCheck, fields...), getFieldsMap); err != nil {
		return nil, nil, err
	}

	whereStr, whereValue := b.renderWhere(where)
	fieldsStr, fieldsAddr, err := b.renderFields(fields, getFieldsMap)
	if err != nil {
		return nil, nil, err
	}
	rows, err := b.query(fmt.Sprintf("SELECT %s FROM `%s` WHERE %s %s", fieldsStr, table, whereStr, extra), whereValue...)
	return rows, fieldsAddr, nil
}

func (b *Model) insertExec(fieldToValueMap structure.StringToObjectMap, getFieldsMap func() structure.StringToObjectMap, table string) (int64, error) {

	//加入默认值
	extraFields := structure.StringToObjectMap{"is_deleted": UnDeleted, "created_at": helper.Now(), "updated_at": helper.Now()}
	for field, val := range extraFields {
		if fieldToValueMap[field] == nil {
			fieldToValueMap[field] = val
		}
	}
	//条件字段校验
	var fieldsToCheck []string
	for field := range fieldToValueMap {
		fieldsToCheck = append(fieldsToCheck, field)
	}
	if err := b.checkFieldValid(fieldsToCheck, getFieldsMap); err != nil {
		return 0, err
	}

	var fields []string
	var alternatives []string
	var values structure.Array
	for field, value := range fieldToValueMap {
		fields = append(fields, fmt.Sprintf("`%s`", field))
		alternatives = append(alternatives, "?")
		values = append(values, value)
	}

	result, err := b.exec(fmt.Sprintf("INSERT INTO `%s`(%s) VALUES (%s)", table, strings.Join(fields, ","), strings.Join(alternatives, ",")), values...)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (b *Model) checkFieldValid(fields []string, getFieldsMap func() structure.StringToObjectMap) error {
	fieldsMap := getFieldsMap()
	for _, field := range fields {
		if addr := fieldsMap[field]; addr == nil {
			return helper.CreateNewError(fmt.Sprintf("invalid key:%s", field))
		}
	}
	return nil
}

func (b *Model) updateExec(fieldToValueMap structure.StringToObjectMap, where structure.StringToObjectMap, getFieldsMap func() structure.StringToObjectMap, table string) (int64, error) {
	//加入默认值
	extraFields := structure.StringToObjectMap{"updated_at": helper.Now()}
	for field, val := range extraFields {
		if fieldToValueMap[field] == nil {
			fieldToValueMap[field] = val
		}
	}

	//条件字段校验
	var fieldsToCheck []string
	for _, object := range []structure.StringToObjectMap{fieldToValueMap, where} {
		for field := range object {
			fieldsToCheck = append(fieldsToCheck, field)
		}
	}
	if err := b.checkFieldValid(fieldsToCheck, getFieldsMap); err != nil {
		return 0, err
	}

	var fields []string
	var values structure.Array
	for field, value := range fieldToValueMap {
		fields = append(fields, fmt.Sprintf("`%s`= ?", field))
		values = append(values, value)
	}

	whereStr, whereValueArr := b.renderWhere(where)
	values = append(values, whereValueArr...)
	result, err := b.exec(fmt.Sprintf("UPDATE `%s` SET %s WHERE %s", table, strings.Join(fields, ","), whereStr), values...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (b *Model) renderFields(fields []string, getFieldsMap func() structure.StringToObjectMap) (string, structure.Array, error) {

	var fieldsToReturn []string
	var addrToReturn structure.Array
	fieldsMap := getFieldsMap()
	if 0 == len(fields) { //没有指定 则取所有字段
		for key := range fieldsMap {
			fields = append(fields, key)
		}
	}

	for _, field := range fields {
		addr := fieldsMap[field]
		if addr == nil {
			return "", nil, helper.CreateNewError("invalid key " + field)
		}

		fieldsToReturn = append(fieldsToReturn, fmt.Sprintf("`%s`", field))
		addrToReturn = append(addrToReturn, addr)
	}

	return strings.Join(fieldsToReturn, ","), addrToReturn, nil
}

func (b *Model) renderWhere(where structure.StringToObjectMap) (string, structure.Array) {
	var whereIndex []string
	var whereValue structure.Array
	whereIndex = append(whereIndex, " 1=1 ")
	for i, v := range where {
		whereIndex = append(whereIndex, fmt.Sprintf(" `%s`= ? ", i))
		whereValue = append(whereValue, v)
	}
	return strings.Join(whereIndex, "and"), whereValue
}
