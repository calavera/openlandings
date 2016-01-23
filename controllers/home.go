package controllers

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	u := c.GetSession("current_user")
	if u != nil {
		c.Redirect("/steps/browse", 302)
	} else {
		c.TplName = "home/index.html"
	}
}
