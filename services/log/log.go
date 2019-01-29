package log

import (
	"MPMS/helper"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func init() {
	logs.Async(1e3)
	logs.SetLogFuncCall(true)
}

func setLogger() {
	_ = logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"logs/%s/%s.log"}`, helper.NowDate(), helper.NowHour()))
}

func Info(f interface{}, v ...interface{}) {
	setLogger()
	logs.Info(f, v)
}

func Err(f interface{}, v ...interface{}) {
	setLogger()
	logs.Error(f, v)
}
