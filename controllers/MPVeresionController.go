package controllers

import (
	"MPMS/helper"
	"MPMS/models"
	"MPMS/routers/uris"
	"MPMS/session"
	"fmt"
)

type MPVersionController struct {
	Controller
}

func (mpv *MPVersionController) Create() {
	req := struct {
		CompanyId int64 `form:"company_id"`
	}{}
	_ = mpv.ParseForm(&req)
	if req.CompanyId == 0 { //没有获取到公司id 那么是用户登陆 则需要获取用户id
		company, err := mpv.getSessionCompanyInfo()
		if err != nil {
			panic("用户没有获取到公司信息！")
		}
		req.CompanyId = company.Id
	}

	mpv.Data["CompanyId"] = req.CompanyId
	mpv.Data["OperateType"] = helper.OperateTypeCreate               //创建
	mpv.Data["ApiUriMiniProgramEdit"] = uris.ApiUriMiniProgramEdit   //编辑接口
	mpv.Data["HtmlUriMiniProgramEdit"] = uris.HtmlUriMiniProgramEdit //编辑页面
	mpv.Data["MiniProgramVersionTypeToNameMap"] = models.MiniProgramVersionTypeToNameMap()
	mpv.RenderHtml("小程序版本创建", "mpv", "mini_program_version/edit/html.tpl", "mini_program_version/edit/css.tpl", "mini_program_version/edit/js.tpl", "")
}

func (mpv *MPVersionController) Edit() {
	req := struct {
		Id int64 `form:"mp_id"`
	}{}
	_ = mpv.ParseForm(&req)

	mpv.Data["MpId"] = req.Id
	mpv.Data["OperateType"] = helper.OperateTypeEdit //创建
	mpv.Data["ApiUriMiniProgramEdit"] = uris.ApiUriMiniProgramEdit
	mpv.Data["MiniProgramVersionTypeToNameMap"] = models.MiniProgramVersionTypeToNameMap()
	mpv.Data["ApiUriMiniProgramList"] = uris.ApiUriMiniProgramList
	mpv.RenderHtml("小程序版本编辑", "mpv", "mini_program_version/edit/html.tpl", "mini_program_version/edit/css.tpl", "mini_program_version/edit/js.tpl", "")
}

func (mpv *MPVersionController) Index() {
	req := struct {
		MpId int64 `form:"mp_id"`
	}{}
	_ = mpv.ParseForm(&req)
	mpv.Data["ApiUriMiniProgramVersionList"] = uris.ApiUriMiniProgramVersionList
	mpv.Data["HtmlUriMiniProgramVersionEdit"] = uris.HtmlUriMiniProgramVersionEdit                                           //编辑页面
	mpv.Data["HtmlUriMiniProgramVersionCreate"] = fmt.Sprintf("%s?mp_id=%d", uris.HtmlUriMiniProgramVersionCreate, req.MpId) //创建页面
	mpv.Data["MpId"] = req.MpId
	mpv.Data["UserType"] = mpv.GetSession(session.UserType)
	mpv.RenderHtml("小程序版本管理", "mpv", "mini_program_version/index/html.tpl", "mini_program_version/index/css.tpl", "mini_program_version/index/js.tpl", "")
}
