package main

import (
	_ "MPMS/routers"
	"MPMS/services/db"
	"MPMS/services/email"
	"github.com/astaxie/beego"
)

func init() {
	//数据库初始化连接池
	go db.InitConPools()
	//处理脚本 写在最前面
	go email.Listen()
}

func main() {
	beego.Run()
}
