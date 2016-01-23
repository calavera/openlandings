package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	u := c.GetSession("current_user")
	if u != nil {
		c.TplName = "home/dashboard.html"
	} else {
		c.TplName = "home/index.html"
	}
}
