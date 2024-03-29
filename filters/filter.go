package filters

import (
	"MPMS/routers/uris"
	"MPMS/services/log"
	"MPMS/session"
	"github.com/astaxie/beego/context"
	"github.com/petermattis/goid"
	"strings"
)

func Before(ctx *context.Context) {
	log.Info("Access", goid.Get(), *ctx.Request)

	isLogin := false
	if uuid := ctx.Input.Session(session.UUID); uuid != nil {
		isLogin = true
	}

	uri := ctx.Request.RequestURI
	can := func(uri string) bool {
		//特殊路由 可以不需要登录
		exceptUris := uris.GetUnCheckLoginUris()
		for _, exceptUri := range exceptUris {
			if uri == exceptUri || strings.Contains(uri, exceptUri+"?") {
				return true
			}
		}
		return false
	}

	//如果未登陆并且路由必须登陆则跳转到登陆
	if !isLogin && !can(uri) {
		//链接写入session 为了登陆成功后跳转
		if strings.ToLower(ctx.Request.Method) == "get" {
			ctx.Output.Session(session.RequestUri, uri)
		}
		ctx.Redirect(302, uris.HtmlUriUserLogin)
	}
}
