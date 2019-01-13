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
	mp.Data["OperateType"] = MPOperateTypeCreate //创建
	mp.Data["MiniProgramTypeToNameMap"] = models.MiniProgramTypeToNameMap()
	mp.RenderHtml("小程序创建", "mp", "mini_program/edit.tpl", "mini_program/css.tpl", "mini_program/scripts.tpl", "")
}
