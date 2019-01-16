package api

import (
	"MPMS/helper"
	"MPMS/models"
	"MPMS/session"
	"MPMS/structure"
	"encoding/json"
	"github.com/astaxie/beego"
)

type Controller struct {
	beego.Controller
}

func (c *Controller) ApiReturn(res structure.Response) {
	c.Data["json"] = res
	c.ServeJSON()
}

func (c *Controller) ParseJsonData(data interface{}) error {
	return json.Unmarshal(c.Ctx.Input.RequestBody, &data)
}

func (c *Controller) getSessionCompanyInfo() (company models.Company, err error) {
	if c.GetSession(session.UserType).(uint8) == models.UserTypeCustomer {
		companyInfoBytes := c.GetSession(session.CompanyInfo).([]byte)
		_ = json.Unmarshal(companyInfoBytes, &company)
		return company, nil
	}
	return company, helper.CreateNewError("该用户为管理员用户")
}
