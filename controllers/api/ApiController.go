package api

import (
	"MPMS/structure"
	"github.com/astaxie/beego"
)

type Controller struct {
	beego.Controller
}

func (c *Controller) ApiReturn(res structure.Response) {
	c.Data["json"] = res
	c.ServeJSON()
}
