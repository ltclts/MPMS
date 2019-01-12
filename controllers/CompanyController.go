package controllers

import "MPMS/routers/uris"

type CompanyController struct {
	Controller
}

func (c *CompanyController) Index() {
	c.Data["UrlGetList"] = uris.ApiUriCompanyList
	c.RenderHtml("账户管理", "company", "company/index.tpl", "company/css.tpl", "company/scripts.tpl", "")
}
