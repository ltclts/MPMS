package controllers

import (
	"MPMS/models"
	"MPMS/routers/uris"
	"MPMS/session"
)

type UserController struct {
	Controller
}

/**
登陆页面
*/
func (u *UserController) Login() {
	u.Data["Title"] = "小程序后台管理系统"
	u.Data["xsrfdata"] = u.XSRFToken()
	u.Data["ApiUriLogin"] = uris.ApiUriUserLogin
	u.TplName = "login/index.tpl"
}

func (u *UserController) InfoChange() {
	u.Data["ApiUriUserGetUserInfo"] = uris.ApiUriUserGetUserInfo
	u.Data["ApiUriUserInfoChange"] = uris.ApiUriUserInfoChange
	u.Data["HtmlUriUserLogin"] = uris.HtmlUriUserLogin
	u.RenderHtml("信息修改", "user", "user/edit/html.tpl", "user/edit/css.tpl", "user/edit/js.tpl", "")
}

func (u *UserController) Get() {
	redirect := uris.HtmlUriCompany
	if u.GetSession(session.UserType).(uint8) == models.UserTypeCustomer {
		redirect = uris.HtmlUriMiniProgram
	}
	u.Ctx.Redirect(302, redirect)
}
