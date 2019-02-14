package job

import "time"

/**
定时脚本 调用方法 间隔时间(单位：s)
*/
func Run(callable func(), seconds float64) {
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
