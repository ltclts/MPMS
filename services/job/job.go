package job

import (
	"github.com/robfig/cron"
	"time"
)

//Spec 取值举例
//每隔5秒执行一次: */5 * * * * ?
//每隔1分钟执行一次: 0 */1 * * * ?
//每天23点执行一次: 0 0 23 * * ?
//每天凌晨1点执行一次: 0 0 1 * * ?
//每月1号凌晨1点执行一次: 0 0 1 1 * ?
//在26分、29分、33分执行一次: 0 26,29,33 * * * ?
//每天的0点、13点、18点、21点都执行一次: 0 0 0,13,18,21 * * ?
type Cron struct {
	Callable func()
	Spec     string
}

type Polling struct {
	Callable func()
	Seconds  float64
}

/**
轮询脚本 调用方法 间隔时间(单位:s)
*/
func pollingRun(callable func(), seconds float64) {
	go func() {
		lastCarriedOutTime := time.Now()
		for true {
			now := time.Now()
			if now.Sub(lastCarriedOutTime).Seconds() >= seconds {
				go callable()
				lastCarriedOutTime = now
			}
			time.Sleep(time.Duration(seconds) * time.Second)
		}
	}()
}

/**
轮询脚本处理
*/
func PollingListRun(pollingList ...Polling) {
	for _, item := range pollingList {
		pollingRun(item.Callable, item.Seconds)
	}
}

/**
  定时脚本
*/
func CronListRun(cronList ...Cron) {
	cronIns := cron.New()
	for _, item := range cronList {
		err := cronIns.AddFunc(item.Spec, item.Callable)
		if err != nil {
			panic(err.Error())
		}
	}

	cronIns.Start()
}
