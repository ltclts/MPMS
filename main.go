package main

import (
	_ "MPMS/routers"
	"MPMS/services/db"
	"MPMS/services/email"
	"MPMS/services/job"
	"MPMS/services/oss"
	"github.com/astaxie/beego"
)

func init() {
	//数据库配置
	go db.InitConfig()
	//处理脚本 写在最前面
	go email.Listen()

	//每天执行一次oss清理服务
	job.CronListRun(job.Cron{Callable: oss.Remove, Spec: "0 0 0 * * ?"})
}

func main() {
	beego.Run()
}
