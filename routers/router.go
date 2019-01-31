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
	beego.InsertFilter("/*", beego.BeforeRouter, filters.Before)

	/**
	路由设定
	*/
	Routes = append(Routes, Route{uris.HtmlUriIndex, &controllers.UserController{}, "get:Get"})
	Routes = append(Routes, Route{uris.HtmlUriUserInfoChange, &controllers.UserController{}, "get:InfoChange"})
	Routes = append(Routes, Route{uris.HtmlUriUserLogin, &controllers.UserController{}, "get:Login"})

	Routes = append(Routes, Route{uris.HtmlUriCompany, &controllers.CompanyController{}, "get:Index"})
	Routes = append(Routes, Route{uris.HtmlUriCompanyCreate, &controllers.CompanyController{}, "get:Create"})
	Routes = append(Routes, Route{uris.HtmlUriCompanyEdit, &controllers.CompanyController{}, "get:Edit"})

	Routes = append(Routes, Route{uris.HtmlUriMiniProgram, &controllers.MPController{}, "get:Index"})
	Routes = append(Routes, Route{uris.HtmlUriMiniProgramCreate, &controllers.MPController{}, "get:Create"})
	Routes = append(Routes, Route{uris.HtmlUriMiniProgramEdit, &controllers.MPController{}, "get:Edit"})

	Routes = append(Routes, Route{uris.HtmlUriMiniProgramVersion, &controllers.MPVersionController{}, "get:Index"})
	Routes = append(Routes, Route{uris.HtmlUriMiniProgramVersionCreate, &controllers.MPVersionController{}, "get:Create"})
	Routes = append(Routes, Route{uris.HtmlUriMiniProgramVersionEdit, &controllers.MPVersionController{}, "get:Edit"})

	Routes = append(Routes, Route{uris.ApiUriUserGetUserInfo, &api.UserApiController{}, "get:GetUserInfo"})
	Routes = append(Routes, Route{uris.ApiUriUserLogin, &api.UserApiController{}, "post:Login"})
	Routes = append(Routes, Route{uris.ApiUriUserLogout, &api.UserApiController{}, "post:Logout"})
	Routes = append(Routes, Route{uris.ApiUriUserGetCheckCode, &api.UserApiController{}, "get:GetCheckCode"})
	Routes = append(Routes, Route{uris.ApiUriUserInfoChange, &api.UserApiController{}, "post:InfoChange"})

	Routes = append(Routes, Route{uris.ApiUriCompanyList, &api.CompanyApiController{}, "post:List"})
	Routes = append(Routes, Route{uris.ApiUriCompanyEdit, &api.CompanyApiController{}, "post:Edit"})
	Routes = append(Routes, Route{uris.ApiUriCompanyGetEditInfo, &api.CompanyApiController{}, "get:GetEditInfo"})
	Routes = append(Routes, Route{uris.ApiUriCompanyUpdateStatus, &api.CompanyApiController{}, "post:UpdateStatus"})

	//小程序创建、编辑接口
	Routes = append(Routes, Route{uris.ApiUriMiniProgramEdit, &api.MPApiController{}, "post:Edit"})
	Routes = append(Routes, Route{uris.ApiUriMiniProgramList, &api.MPApiController{}, "post:List"})

	Routes = append(Routes, Route{uris.ApiUriMiniProgramVersionList, &api.MPVersionApiController{}, "post:List"})
	Routes = append(Routes, Route{uris.ApiUriMiniProgramVersionUpload, &api.MPVersionApiController{}, "post:Upload"})
	Routes = append(Routes, Route{uris.ApiUriMiniProgramVersionGet, &api.MPVersionApiController{}, "get:Get"})
	Routes = append(Routes, Route{uris.ApiUriMiniProgramVersionEdit, &api.MPVersionApiController{}, "post:Edit"})

	/**
	外部调用接口
	*/
	Routes = append(Routes, Route{uris.ApiUriMpOutInfoGet, &api.MPOutApiController{}, "post:RequestInfo"})

	Routes = append(Routes, Route{uris.ApiUriHelperRefreshDBConPools, &api.HelperApiController{}, "get:RefreshDBConPools"})

	for _, route := range Routes {
		beego.Router(route.Uri, route.ControllerInterface, route.Method)
	}
}
