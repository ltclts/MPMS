package main

import (
	_ "MPMS/routers"
	"MPMS/services/db"
	"MPMS/services/email"
	"github.com/astaxie/beego"
)

func init() {
	//数据库配置
	go db.InitConfig()
	//处理脚本 写在最前面
	go email.Listen()
}

func main() {
	beego.Run()
}
