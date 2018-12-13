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
type model struct {
	IsDeleted uint8
	CreatorId uint
	CreatedAt string
	UpdatedAt string
}

const UNDELETED = 0
const DELETED = 1

var db *sql.DB
var tx *sql.Tx        //事务使用
var userTrans = false //是否开启事务

/**
连接数据库（使用单例模式）
*/
func (b *model) InitDB() (*sql.DB, error) {
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
func (b *model) StartTrans() (*sql.Tx, error) {
	var err error
	if db == nil {
		db, err = b.InitDB()
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
func (b *model) Rollback() error {
	err := tx.Rollback()
	if err == nil {
		userTrans = false
	}
	return err
}

/**
事务回滚
*/
func (b *model) Commit() error {
	err := tx.Commit()
	if err == nil {
		userTrans = false
	}
	return err
}

/**
查询方法
*/
func (b *model) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	if db == nil {
		if _, err := b.InitDB(); err != nil {
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
func (b *model) Exec(sql string, args ...interface{}) (sql.Result, error) {

	if db == nil {
		if _, err := b.InitDB(); err != nil {
			return nil, err
		}
	}

	if userTrans {
		smt, _ := tx.Prepare(sql)
		return smt.Exec(args...)
	}

	return db.Exec(sql, args...)
}

func (b *model) renderFields(fields []string, getFieldsMap func() structure.Map) (string, structure.Array, error) {

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
			return "", nil, helper.ThrowNewError("invalid key " + field)
		}

		fieldsToReturn = append(fieldsToReturn, "`"+field+"`")
		addrToReturn = append(addrToReturn, addr)
	}

	return strings.Join(fieldsToReturn, ","), addrToReturn, nil
}

func (b *model) renderWhere(where structure.Map) (string, structure.Array) {
	var whereIndex []string
	var whereValue structure.Array
	whereIndex = append(whereIndex, " 1=1 ")
	for i, v := range where {
		whereIndex = append(whereIndex, " `"+i+"`= ? ")
		whereValue = append(whereValue, v)
	}
	return strings.Join(whereIndex, "and"), whereValue
}
