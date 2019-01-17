package controllers

import (
	"MPMS/helper"
	"MPMS/routers/uris"
)

type CompanyController struct {
	Controller
}

func (c *CompanyController) Index() {
	company, err := c.getSessionCompanyInfo()
	if nil == err {
		c.Data["CompanyId"] = company.Id
	}
	c.Data["UrlGetList"] = uris.ApiUriCompanyList
	c.Data["HtmlUriCompanyCreate"] = uris.HtmlUriCompanyCreate //创建页面
	c.Data["HtmlUriCompanyEdit"] = uris.HtmlUriCompanyEdit     //编辑页面
	c.Data["HtmlUriMiniProgramCreate"] = uris.HtmlUriMiniProgramCreate
	c.RenderHtml("账户管理", "company", "company/index/html.tpl", "company/index/css.tpl", "company/index/js.tpl", "")
}

func (c *CompanyController) Create() {
	c.Data["OperateType"] = helper.OperateTypeCreate               //创建
	c.Data["HtmlUriCompanyEdit"] = uris.HtmlUriCompanyEdit         //编辑页面
	c.Data["ApiUriCompanyEdit"] = uris.ApiUriCompanyEdit           //编辑接口
	c.Data["ApiUriUserGetCheckCode"] = uris.ApiUriUserGetCheckCode //获取验证码

	c.RenderHtml("公司创建", "company", "company/edit/html.tpl", "company/edit/css.tpl", "company/edit/js.tpl", "")
}

func (c *CompanyController) Edit() {

	req := struct {
		CompanyId int64 `form:"company_id"`
	}{}
	_ = c.ParseForm(&req)
	if req.CompanyId == 0 { //没有获取到公司id 那么是用户登陆 则需要获取用户id
		company, err := c.getSessionCompanyInfo()
		if err != nil {
			panic("用户没有获取到公司信息！")
		}
		req.CompanyId = company.Id
	}

	c.Data["Id"] = req.CompanyId
	c.Data["ApiUriCompanyGetEditInfo"] = uris.ApiUriCompanyGetEditInfo
	c.Data["OperateType"] = helper.OperateTypeEdit                 //编辑
	c.Data["ApiUriCompanyEdit"] = uris.ApiUriCompanyEdit           //编辑接口
	c.Data["ApiUriUserGetCheckCode"] = uris.ApiUriUserGetCheckCode //获取验证码
	c.RenderHtml("公司编辑", "company", "company/edit/html.tpl", "company/edit/css.tpl", "company/edit/js.tpl", "")
}
