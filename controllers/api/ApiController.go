package api

import (
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
