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
	TestSql                = `select 1`
	QueryMaxCountSql       = `show variables like 'max_connections'`
	QueryMaxWaitTimeOutSql = `show variables like 'wait_timeout'`
)

var ConCount = 0 //当前DB连接数
var UerToConMap map[int64]*Con

var MaxConCount = 200   //默认最大连接数
var MaxWaitTimeOut = 30 //超时时间

//当队列用
var ConPools = make(chan *Con, MaxConCount)

type Con struct {
	db           *sql.DB //数据库连接
	tx           *sql.Tx //事务连接
	useTrans     bool    //是否使用事务
	lastLiveTime time.Time
}

func InitConPools() {
	_ = QueryMaxCount()
	_ = QueryMaxWaitTimeOut()

	//初始化 1/10 - 当前已连接数的连接
	initCount := MaxConCount/10 - ConCount
	for i := 0; i < initCount; i++ {
		con, err := initCon()
		if err != nil {
			log.Err("初始化数据库连接失败", err.Error())
			return
		}
		ConPools <- con
	}
	log.Info("数据库连接初始化结束", ConCount, MaxConCount, MaxWaitTimeOut, len(ConPools), cap(ConPools))
}

//测试连通
func (con *Con) Test() error {
	_, err := con.db.Query(TestSql)
	if err != nil {
		return err
	}
	return nil
}

func QueryMaxCount() error {
	rows, err := Query(QueryMaxCountSql)
	if err != nil {
		log.Err("获取数据库最大连接数失败", err.Error(), QueryMaxCountSql)
		return err
	}
	defer func() {
		_ = rows.Close()
	}()

	var name string
	var val int
	for rows.Next() {
		err = rows.Scan(&name, &val)
		if err != nil {
			log.Err("读取数据库最大连接数失败", err.Error(), QueryMaxCountSql)
			return err
		}
	}

	MaxConCount = val / 2

	//重新创建连接池
	conPoolsCopy := make(chan *Con, MaxConCount)
	for len(ConPools) > 0 {
		con := <-ConPools
		conPoolsCopy <- con
	}
	ConPools = conPoolsCopy
	return nil
}

func QueryMaxWaitTimeOut() error {
	rows, err := Query(QueryMaxWaitTimeOutSql)
	if err != nil {
		log.Err("获取数据库最大超时时间失败", err.Error(), QueryMaxWaitTimeOutSql)
		return err
	}
	defer func() {
		_ = rows.Close()
	}()
	var name string
	for rows.Next() {
		err = rows.Scan(&name, &MaxWaitTimeOut)
		if err != nil {
			log.Err("读取数据库最大超时时间失败", err.Error(), QueryMaxWaitTimeOutSql)
			return err
		}
	}
	return nil
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
				if int(time.Now().Sub(con.lastLiveTime).Seconds()) > MaxWaitTimeOut {
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

		con, err = initCon()
		if err != nil {
			return con, err
		}

		if UerToConMap == nil {
			UerToConMap = map[int64]*Con{}
		}
		UerToConMap[unique] = con
		return con, nil
	}
}

func initCon() (con *Con, err error) {
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
	return con, nil
}

//释放
func Release() error {
	con, err := connect()
	if err != nil {
		return err
	}

	delete(UerToConMap, goid.Get())
	ConPools <- con
	return nil
}

func (con *Con) release() error {
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
		_ = con.release() //释放占用
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
			_ = con.release() //释放占用
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
	_ = con.release()
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
	_ = con.release()
	return result, err
}
