package controllers

import "MPMS/models"

type MPController struct {
	Controller
}

const (
	MPOperateTypeCreate = 0
	MPOperateTypeEdit   = 1
)

func (mp *MPController) Create() {
	param := struct {
		CompanyId string `form:"company_id"`
	}{}
	_ = mp.ParseForm(&param)
	mp.Data["CompanyId"] = param.CompanyId
	mp.Data["OperateType"] = MPOperateTypeCreate //创建
	mp.Data["MiniProgramTypeToNameMap"] = models.MiniProgramTypeToNameMap()
	mp.RenderHtml("小程序创建", "mp", "mini_program/edit.tpl", "mini_program/css.tpl", "mini_program/scripts.tpl", "")
}
