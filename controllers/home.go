package controllers

import (
	"github.com/astaxie/beego"
	"github.com/markbates/goth"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	u := c.GetSession("current_user")
	if u != nil {
		c.userDashboard(u.(*goth.User))
	} else {
		c.TplName = "home/index.html"
	}
}

func (c *HomeController) userDashboard(user *goth.User) {
	c.TplName = "home/dashboard.tmpl"
}
