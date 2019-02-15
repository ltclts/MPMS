package models

import (
	"MPMS/helper"
	"MPMS/services/db"
	"MPMS/structure"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"reflect"
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
	ConfigTableName             = "config"
	UserTableName               = "user"
	MenuTableName               = "menu"
	FlowTableName               = "flow"
	RelationTableName           = "relation"
	CompanyTableName            = "company"
	MiniProgramTableName        = "mini_program"
	MiniProgramVersionTableName = "mini_program_version"
	ResourceTableName           = "resource"

	//删除标志
	UnDeleted = 0
	Deleted   = 1
	Removed   = 2
)

/**
开启事务
*/
func (b *Model) StartTrans() (*sql.Tx, error) {
	return db.StartTrans()
}

/**
事务回滚
*/
func (b *Model) Rollback() error {
	return db.Rollback()
}

/**
事务回滚
*/
func (b *Model) Commit() error {
	return db.Commit()
}

/**
查询方法
*/
func (b *Model) query(sql string, args ...interface{}) (rows *sql.Rows, err error) {
	return db.Query(sql, args...)
}

/**
执行方法
*/
func (b *Model) exec(sql string, args ...interface{}) (result sql.Result, err error) {
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
	return rows, fieldsAddr, err
}

func (b *Model) count(getFieldsMap func() structure.StringToObjectMap, where structure.StringToObjectMap, table string) (int64, error) {
	var count int64 = 0
	//条件字段校验
	var fieldsToCheck []string
	for field := range where {
		fieldsToCheck = append(fieldsToCheck, field)
	}
	if err := b.checkFieldValid(fieldsToCheck, getFieldsMap); err != nil {
		return count, err
	}
	whereStr, whereValue := b.renderWhere(where)
	rows, err := b.query(fmt.Sprintf("SELECT count(id) as count FROM `%s` WHERE %s", table, whereStr), whereValue...)
	if err != nil {
		return count, err
	}

	defer func() { _ = rows.Close() }()
	for rows.Next() {
		err = rows.Scan(&count)
		return count, err
	}
	return count, nil
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
		switch v.(type) {
		case structure.Array:
			arr := reflect.ValueOf(v).Interface().(structure.Array)
			if len(arr) != 2 {
				panic("require two params")
			}
			whereIndex = append(whereIndex, fmt.Sprintf(" `%s` %s ? ", i, arr[0]))
			whereValue = append(whereValue, arr[1])
		default:
			whereIndex = append(whereIndex, fmt.Sprintf(" `%s`= ? ", i))
			whereValue = append(whereValue, v)
		}

	}
	return strings.Join(whereIndex, "and"), whereValue
}

func (b *Model) renderWhereDirectly(where structure.StringToObjectMap) (string, structure.Array) {
	var whereIndex []string
	var whereValue structure.Array
	whereIndex = append(whereIndex, " 1=1 ")
	for i, v := range where {
		switch v.(type) {
		case structure.Array:
			arr := reflect.ValueOf(v).Interface().(structure.Array)
			if len(arr) != 2 {
				panic("require two params")
			}
			whereIndex = append(whereIndex, fmt.Sprintf(" %s %s ? ", i, arr[0]))
			whereValue = append(whereValue, arr[1])
		default:
			whereIndex = append(whereIndex, fmt.Sprintf(" %s= ? ", i))
			whereValue = append(whereValue, v)
		}
	}
	return strings.Join(whereIndex, "and"), whereValue
}
