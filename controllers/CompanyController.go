package controllers

import (
	"MPMS/helper"
	"MPMS/routers/uris"
)

type CompanyController struct {
	Controller
}

func (mp *CompanyController) Index() {
	company, err := mp.getSessionCompanyInfo()
	if nil == err {
		mp.Data["CompanyId"] = company.Id
	}
	mp.Data["UrlGetList"] = uris.ApiUriCompanyList
	mp.Data["HtmlUriCompanyCreate"] = uris.HtmlUriCompanyCreate //创建页面
	mp.Data["HtmlUriCompanyEdit"] = uris.HtmlUriCompanyEdit     //编辑页面
	mp.Data["HtmlUriMiniProgramCreate"] = uris.HtmlUriMiniProgramCreate
	mp.Data["ApiUriCompanyUpdateStatus"] = uris.ApiUriCompanyUpdateStatus //状态变更接口
	mp.RenderHtml("账户管理", "company", "company/index/html.tpl", "company/index/css.tpl", "company/index/js.tpl", "")
}

func (mp *CompanyController) Create() {
	mp.Data["OperateType"] = helper.OperateTypeCreate               //创建
	mp.Data["HtmlUriCompanyEdit"] = uris.HtmlUriCompanyEdit         //编辑页面
	mp.Data["ApiUriCompanyEdit"] = uris.ApiUriCompanyEdit           //编辑接口
	mp.Data["ApiUriUserGetCheckCode"] = uris.ApiUriUserGetCheckCode //获取验证码

	mp.RenderHtml("公司创建", "company", "company/edit/html.tpl", "company/edit/css.tpl", "company/edit/js.tpl", "")
}

func (mp *CompanyController) Edit() {

	req := struct {
		CompanyId int64 `form:"company_id"`
	}{}
	_ = mp.ParseForm(&req)
	if req.CompanyId == 0 { //没有获取到公司id 那么是用户登陆 则需要获取用户id
		company, err := mp.getSessionCompanyInfo()
		if err != nil {
			panic("用户没有获取到公司信息！")
		}
		req.CompanyId = company.Id
	}

	mp.Data["Id"] = req.CompanyId
	mp.Data["ApiUriCompanyGetEditInfo"] = uris.ApiUriCompanyGetEditInfo
	mp.Data["OperateType"] = helper.OperateTypeEdit                 //编辑
	mp.Data["ApiUriCompanyEdit"] = uris.ApiUriCompanyEdit           //编辑接口
	mp.Data["ApiUriUserGetCheckCode"] = uris.ApiUriUserGetCheckCode //获取验证码
	mp.RenderHtml("公司编辑", "company", "company/edit/html.tpl", "company/edit/css.tpl", "company/edit/js.tpl", "")
}
