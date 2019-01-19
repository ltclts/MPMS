package controllers

import (
	"MPMS/helper"
	"MPMS/models"
	"MPMS/routers/uris"
	"MPMS/session"
)

type MPController struct {
	Controller
}

func (mp *MPController) Create() {
	mpCreateReq := struct {
		CompanyId int64 `form:"company_id"`
	}{}
	_ = mp.ParseForm(&mpCreateReq)
	if mpCreateReq.CompanyId == 0 { //没有获取到公司id 那么是用户登陆 则需要获取用户id
		company, err := mp.getSessionCompanyInfo()
		if err != nil {
			panic("用户没有获取到公司信息！")
		}
		mpCreateReq.CompanyId = company.Id
	}

	mp.Data["CompanyId"] = mpCreateReq.CompanyId
	mp.Data["OperateType"] = helper.OperateTypeCreate               //创建
	mp.Data["ApiUriMiniProgramEdit"] = uris.ApiUriMiniProgramEdit   //编辑接口
	mp.Data["HtmlUriMiniProgramEdit"] = uris.HtmlUriMiniProgramEdit //编辑页面
	mp.Data["MiniProgramVersionTypeToNameMap"] = models.MiniProgramVersionTypeToNameMap()
	mp.RenderHtml("小程序创建", "mp", "mini_program/edit/html.tpl", "mini_program/edit/css.tpl", "mini_program/edit/js.tpl", "")
}

func (mp *MPController) Edit() {
	req := struct {
		Id int64 `form:"mp_id"`
	}{}
	_ = mp.ParseForm(&req)

	mp.Data["Id"] = req.Id
	mp.Data["OperateType"] = helper.OperateTypeEdit //创建
	mp.Data["ApiUriMiniProgramEdit"] = uris.ApiUriMiniProgramEdit
	mp.Data["MiniProgramVersionTypeToNameMap"] = models.MiniProgramVersionTypeToNameMap()
	mp.Data["ApiUriMiniProgramList"] = uris.ApiUriMiniProgramList
	mp.RenderHtml("小程序编辑", "mp", "mini_program/edit/html.tpl", "mini_program/edit/css.tpl", "mini_program/edit/js.tpl", "")
}

func (mp *MPController) Index() {
	req := struct {
		CompanyId int64 `form:"company_id"`
	}{}
	_ = mp.ParseForm(&req)
	if req.CompanyId == 0 && mp.GetSession(session.UserType).(uint8) == models.UserTypeCustomer {
		company, err := mp.getSessionCompanyInfo()
		if err != nil {
			panic("用户没有获取到公司信息！")
		}
		req.CompanyId = company.Id
	}
	mp.Data["CompanyId"] = req.CompanyId
	mp.Data["ApiUriMiniProgramList"] = uris.ApiUriMiniProgramList
	mp.Data["HtmlUriMiniProgramEdit"] = uris.HtmlUriMiniProgramEdit       //编辑页面
	mp.Data["HtmlUriMiniProgramCreate"] = uris.HtmlUriMiniProgramCreate   //创建页面
	mp.Data["HtmlUriMiniProgramVersion"] = uris.HtmlUriMiniProgramVersion //版本管理页面
	mp.Data["UserType"] = mp.GetSession(session.UserType)
	mp.RenderHtml("小程序管理", "mp", "mini_program/index/html.tpl", "mini_program/index/css.tpl", "mini_program/index/js.tpl", "")
}
