package log

import (
	"MPMS/helper"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func init() {
	logs.EnableFuncCallDepth(true)
	logs.Async(1e3)
	_ = logs.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"logs/%s/%s.log"}`, helper.NowDate(), helper.NowHour()))
}

func Info(f interface{}, v ...interface{}) {
	logs.Info(f, v)
}

func Err(f interface{}, v ...interface{}) {
	logs.Error(f, v)
}
