package db

import (
	"MPMS/helper"
	"MPMS/services/log"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/petermattis/goid"
	"runtime/debug"
	"time"
)

const (
	MaxConCount = 300 //最大连接数
	WaitTimeOut = 300 //超时时间 todo 在开始位置读取数据库超时时间
)

var ConCount = 0 //当前DB连接数

var UerToConMap map[int64]*Con

//当队列用
var ConPools = make(chan *Con, MaxConCount)

type Con struct {
	db           *sql.DB //数据库连接
	tx           *sql.Tx //事务连接
	useTrans     bool    //是否使用事务
	lastLiveTime time.Time
}

func connect() (con *Con, err error) {
	unique := goid.Get()
	if con = UerToConMap[unique]; con != nil {
		con.lastLiveTime = time.Now()
		return con, nil
	} else {
		//判断连接是否超时
		for len(ConPools) > 0 {
			if con = <-ConPools; con != nil {
				if time.Now().Sub(con.lastLiveTime).Seconds() > WaitTimeOut {
					_ = con.db.Close() //关闭连接
					ConCount--
					continue
				}

				con.lastLiveTime = time.Now()
				UerToConMap[unique] = con
				return con, nil
			}
		}

		if ConCount > MaxConCount {
			log.Err("数据库超出最大连接数", ConCount, MaxConCount)
			return con, helper.CreateNewError("max connections")
		}

		ConCount++
		db, err := initDB()
		if err != nil {
			ConCount--
			return con, err
		}
		con = new(Con)
		con.db = db
		con.lastLiveTime = time.Now()
		con.useTrans = false
		if UerToConMap == nil {
			UerToConMap = map[int64]*Con{}
		}
		UerToConMap[unique] = con
		return con, nil
	}
}

//测试连通
func (con *Con) Test() error {
	_, err := con.db.Query("select 1")
	if err != nil {
		return err
	}
	return nil
}

//释放
func release() error {
	con, err := connect()
	if err != nil {
		return err
	}

	delete(UerToConMap, goid.Get())
	ConPools <- con
	return nil
}

/**
连接数据库（使用单例模式）
*/
func initDB() (db *sql.DB, err error) {
	if db != nil {
		return db, err
	}
	db, err = sql.Open(beego.AppConfig.String("DBDriverName"), beego.AppConfig.String("DBDataSourceName"))
	if err != nil {
		return db, err
	}
	return db, err
}

/**
开启事务
*/
func StartTrans() (*sql.Tx, error) {
	var tx *sql.Tx
	con, err := connect()
	if err != nil {
		return tx, err
	}

	if con.useTrans {
		return con.tx, err
	}

	if tx, err = con.db.Begin(); err != nil {
		return tx, err
	}
	con.useTrans = true
	con.tx = tx
	return tx, err
}

/**
事务回滚
*/
func Rollback() error {
	con, err := connect()
	if err != nil {
		return err
	}

	if !con.useTrans {
		return nil
	}

	err = con.tx.Rollback()
	if err == nil {
		con.useTrans = false
		con.tx = nil
		_ = release() //释放占用
	}

	return err
}

/**
事务回滚
*/
func Commit() error {
	con, err := connect()
	if err != nil {
		return err
	}

	if con.useTrans {
		err := con.tx.Commit()
		if err == nil {
			con.useTrans = false
			con.tx = nil
			_ = release() //释放占用
		}
	}

	return err
}

/**
查询方法
*/
func Query(sql string, args ...interface{}) (rows *sql.Rows, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = helper.CreateNewError(fmt.Sprintf("%s", r))
			debug.PrintStack()
		}
	}()

	con, err := connect()
	if err != nil {
		return rows, err
	}

	if con.useTrans {
		smt, _ := con.tx.Prepare(sql)
		return smt.Query(args...)
	}
	rows, err = con.db.Query(sql, args...)
	_ = release()
	return rows, err
}

/**
执行方法
*/
func Exec(sql string, args ...interface{}) (result sql.Result, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = helper.CreateNewError(fmt.Sprintf("%s", r))
			debug.PrintStack()
		}
	}()
	con, err := connect()
	if err != nil {
		return result, err
	}

	if con.useTrans {
		smt, _ := con.tx.Prepare(sql)
		return smt.Exec(args...)
	}

	result, err = con.db.Exec(sql, args...)
	_ = release()
	return result, err
}
