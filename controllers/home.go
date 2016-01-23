package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	_, ok := c.Ctx.Input.Session("current-user").(string)
	if ok {
		c.TplName = "home/dashboard.html"
	} else {
		c.TplName = "home/index.html"
	}
}
