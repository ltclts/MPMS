package routers

import (
	"MPMS/controllers"
	"MPMS/controllers/api"
	"MPMS/filters"
	"MPMS/routers/uris"
	"github.com/astaxie/beego"
)

func init() {

	/**
	中间件
	*/
	beego.InsertFilter("/*", beego.BeforeRouter, filters.FilterUser)

	/**
	主页面
	*/
	beego.Router(uris.HtmlUriIndex, &controllers.IndexController{})

	/**
	页面
	*/
	beego.Router(uris.HtmlUriLogin, &controllers.UserController{}, "get:Login")

	/**
	接口
	*/
	beego.Router(uris.ApiUriLogin, &api.UserApiController{}, "post:Login")

}
