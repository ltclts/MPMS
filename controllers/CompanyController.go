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
	c.Data["HtmlUriCompanyCreate"] = uris.HtmlUriCompanyCreate //编辑页面
	c.Data["HtmlUriMiniProgramCreate"] = uris.HtmlUriMiniProgramCreate
	c.RenderHtml("账户管理", "company", "company/index/html.tpl", "company/index/css.tpl", "company/index/js.tpl", "")
}

func (c *CompanyController) Create() {

	c.Data["OperateType"] = helper.OperateTypeCreate       //创建
	c.Data["HtmlUriCompanyEdit"] = uris.HtmlUriCompanyEdit //编辑页面
	c.Data["ApiUriCompanyEdit"] = uris.ApiUriCompanyEdit   //编辑接口
	c.RenderHtml("公司创建", "company", "company/edit/html.tpl", "company/edit/css.tpl", "company/edit/js.tpl", "")
}

func (c *CompanyController) Edit() {
	company, err := c.getSessionCompanyInfo()
	if nil == err {
		c.Data["CompanyId"] = company.Id
	}
	c.Data["UrlGetList"] = uris.ApiUriCompanyList
	c.Data["HtmlUriMiniProgramCreate"] = uris.HtmlUriMiniProgramCreate
	c.RenderHtml("账户管理", "company", "company/index/html.tpl", "company/index/css.tpl", "company/index/js.tpl", "")
}
