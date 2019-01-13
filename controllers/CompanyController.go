package controllers

import (
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
	c.Data["HtmlUriMiniProgramCreate"] = uris.HtmlUriMiniProgramCreate
	c.RenderHtml("账户管理", "company", "company/index.tpl", "company/css.tpl", "company/scripts.tpl", "")
}
