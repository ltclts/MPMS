package filters

import (
	"MPMS/session"
	"fmt"
	"github.com/astaxie/beego/context"
	"strings"
)

func FilterUser(ctx *context.Context) {
	ok := false
	if uuid := ctx.Input.Session(session.UUID); uuid != nil {
		ok = true
	}

	uri := ctx.Request.RequestURI
	can := func(uri string) bool {
		fmt.Println(uri)
		//特殊路由 可以不需要登录
		exceptUris := []string{"/user/login", "/api/user/login"}
		for _, exceptUri := range exceptUris {
			if uri == exceptUri || strings.Contains(uri, exceptUri+"?") {
				return true
			}
		}
		return false
	}

	if !ok && !can(uri) {
		//链接写入session 为了登陆成功后跳转
		if strings.ToLower(ctx.Request.Method) == "get" {
			ctx.Output.Session(session.RequestUri, uri)
		}
		ctx.Redirect(302, "/user/login")
	}
}
