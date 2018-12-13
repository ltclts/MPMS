package routers

import (
	"MPMS/controllers"
	"MPMS/controllers/api"
	"MPMS/filters"
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
	beego.Router("/", &controllers.IndexController{})

	/**
	页面
	*/
	beego.Router("/user/login", &controllers.UserController{}, "get:Login")

	/**
	接口
	*/
	beego.Router("/api/user/login", &api.UserApiController{}, "post:Login")

}
