package main

import (
	_ "MPMS/routers"
	"MPMS/services/email"
	"github.com/astaxie/beego"
)

func init() {
	//处理脚本 写在最前面
	go email.Listen()
}

func main() {
	beego.Run()
}
