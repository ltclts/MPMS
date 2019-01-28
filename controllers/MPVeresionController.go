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
		MpId int64 `form:"mp_id"`
	}{}
	_ = mpv.ParseForm(&req)
	if req.MpId == 0 { //创建必须带着小程序id 否则不允许创建
		mpv.Data["Error"] = "参数错误，请退出页面重新进入"
	}

	mpv.Data["MpId"] = req.MpId
	mpv.Data["OperateType"] = helper.OperateTypeCreate //创建
	mpv.Data["ApiUriMiniProgramVersionEdit"] = uris.ApiUriMiniProgramVersionEdit
	mpv.Data["HtmlUriMiniProgramVersionEdit"] = uris.HtmlUriMiniProgramVersionEdit
	mpv.Data["ApiUriMiniProgramVersionUpload"] = uris.ApiUriMiniProgramVersionUpload
	mpv.Data["MiniProgramVersionTypeToNameMap"] = models.MiniProgramVersionTypeToNameMap()
	mpv.RenderHtml("小程序版本创建", "mpv", "mini_program_version/edit/html.tpl", "mini_program_version/edit/css.tpl", "mini_program_version/edit/js.tpl", "")
}

func (mpv *MPVersionController) Edit() {
	req := struct {
		Id int64 `form:"mini_program_version_id"`
	}{}
	_ = mpv.ParseForm(&req)

	mpv.Data["Id"] = req.Id
	mpv.Data["OperateType"] = helper.OperateTypeEdit //创建
	mpv.Data["ApiUriMiniProgramVersionEdit"] = uris.ApiUriMiniProgramVersionEdit
	mpv.Data["ApiUriMiniProgramVersionGet"] = uris.ApiUriMiniProgramVersionGet
	mpv.Data["ApiUriMiniProgramVersionUpload"] = uris.ApiUriMiniProgramVersionUpload
	mpv.Data["MiniProgramVersionTypeToNameMap"] = models.MiniProgramVersionTypeToNameMap()
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
