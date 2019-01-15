package routers

import (
	"MPMS/controllers"
	"MPMS/controllers/api"
	"MPMS/filters"
	"MPMS/routers/uris"
	"github.com/astaxie/beego"
)

type Route struct {
	Uri                 string
	ControllerInterface beego.ControllerInterface
	Method              string
}

var Routes []Route

func init() {

	/**
	中间件
	*/
	beego.InsertFilter("/*", beego.BeforeRouter, filters.FilterUser)

	/**
	路由设定
	*/
	Routes = append(Routes, Route{uris.HtmlUriIndex, &controllers.IndexController{}, "get:Get"})
	Routes = append(Routes, Route{uris.HtmlUriLogin, &controllers.UserController{}, "get:Login"})
	Routes = append(Routes, Route{uris.ApiUriLogin, &api.UserApiController{}, "post:Login"})
	Routes = append(Routes, Route{uris.ApiUriLogout, &api.UserApiController{}, "post:Logout"})

	Routes = append(Routes, Route{uris.HtmlUriCompany, &controllers.CompanyController{}, "get:Index"})
	Routes = append(Routes, Route{uris.ApiUriCompanyList, &api.CompanyApiController{}, "post:List"})
	Routes = append(Routes, Route{uris.HtmlUriMiniProgramCreate, &controllers.MPController{}, "get:Create"})
	Routes = append(Routes, Route{uris.HtmlUriMiniProgramEdit, &controllers.MPController{}, "get:Edit"})

	//小程序创建、编辑接口
	Routes = append(Routes, Route{uris.ApiUriMiniProgramEdit, &api.MPApiController{}, "post:Edit"})

	/**
	外部调用接口
	*/
	Routes = append(Routes, Route{uris.ApiUriMpOutPageConfigRequest, &api.MPOutApiController{}, "post:PageConfigRequest"})

	for _, route := range Routes {
		beego.Router(route.Uri, route.ControllerInterface, route.Method)
	}
}
