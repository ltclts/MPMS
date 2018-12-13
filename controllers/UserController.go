package controllers

import "MPMS/routers/uris"

type UserController struct {
	Controller
}

/**
登陆页面
*/
func (u *UserController) Login() {
	u.Data["Title"] = "用户登录"
	u.Data["xsrfdata"] = u.XSRFToken()
	u.Data["ApiUriLogin"] = uris.ApiUriLogin
	u.TplName = "login/index.tpl"
}
