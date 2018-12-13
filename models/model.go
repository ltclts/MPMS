package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
)

/**
基础模型 用于写通用方法
@author cyu 2018-4-28 16:29:54
*/
type model struct {
	CreatorId uint
	CreatedAt string
	UpdatedAt string
}

var db *sql.DB
var tx *sql.Tx        //事务使用
var userTrans = false //是否开启事务

/**
连接数据库（使用单例模式）
@author cyu 2018-4-28 17:13:33
*/
func (b *model) InitDB() (*sql.DB, error) {
	var err error
	if db != nil {
		return db, err
	}
	fmt.Println("init -- db")
	db, err = sql.Open(beego.AppConfig.String("driverName"), beego.AppConfig.String("dataSourceName"))
	if err != nil {
		return db, err
	}
	return db, err
}

/**
开启事务
@author cyu 2018-4-28 17:27:22
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
@author cyu 2018-4-28 17:27:22
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
@author cyu 2018-4-28 17:27:22
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
@author cyu 2018-4-29 17:49:08
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
@author cyu 2018-4-29 17:49:08
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
