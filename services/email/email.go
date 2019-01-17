package email

import (
	"MPMS/services/log"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/go-gomail/gomail"
)

type InfoToSend interface {
	Send()
}

type RegisterEmail struct {
	Tos  []To
	Code string
}

type NoticePasswordEmail struct {
	Tos      []To
	Password string
}

type To struct {
	Addr string
	Name string
}

var dialer *gomail.Dialer
var message *gomail.Message
var msgChanList = make(chan InfoToSend, 500)

const (
	RegisterMsgTemplate       = `您的注册验证码为：<font color="red"><I>%s</I></font>，请将该验证码告知工作人员。`
	NoticePasswordMsgTemplate = `您的登陆密码为：<font color="red"><I>%s</I></font>，请勿将密码告知他人！`
)

func init() {
	//发件人配置
	message = gomail.NewMessage()
	message.SetAddressHeader("From", beego.AppConfig.String("email.fromaddr"), beego.AppConfig.String("email.fromname"))

	// 发送邮件服务器、端口、发件人账号、发件人密码
	port, _ := beego.AppConfig.Int("email.port")
	dialer = gomail.NewDialer(beego.AppConfig.String("email.host"), port, beego.AppConfig.String("email.username"), beego.AppConfig.String("email.password"))
}

func SetMsg(msg InfoToSend) {
	log.Info("写入邮件消息队列", msg)
	msgChanList <- msg
}

func (e RegisterEmail) Send() {
	var toStrArr []string
	for _, to := range e.Tos {
		toStrArr = append(toStrArr, message.FormatAddress(to.Addr, to.Name))
	}
	message.SetHeader("To", toStrArr...)
	message.SetHeader("Subject", "注册验证码")
	message.SetBody("text/html", fmt.Sprintf(RegisterMsgTemplate, e.Code)) // 正文
	if err := dialer.DialAndSend(message); err != nil {
		log.Err("发送邮件失败", err)
		return
	}
	log.Info("发送邮件成功", e)
}

func (n NoticePasswordEmail) Send() {
	var toStrArr []string
	for _, to := range n.Tos {
		toStrArr = append(toStrArr, message.FormatAddress(to.Addr, to.Name))
	}
	message.SetHeader("To", toStrArr...)
	message.SetHeader("Subject", "密码通知")
	message.SetBody("text/html", fmt.Sprintf(NoticePasswordMsgTemplate, n.Password)) // 正文
	if err := dialer.DialAndSend(message); err != nil {
		log.Err("发送邮件失败", err)
		return
	}
	log.Info("发送邮件成功", n)
}

func Listen() {
	log.Info("发送邮件服务监听开启")
	var item InfoToSend
	for {
		select {
		case item = <-msgChanList:
			log.Info("获取到将要被发送的邮件", item)
			item.Send()
		}
	}
}
