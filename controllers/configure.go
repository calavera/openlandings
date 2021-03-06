package controllers

import (
	"github.com/astaxie/beego"
	"github.com/calavera/openlandings/github"
	"github.com/calavera/openlandings/models"
	"github.com/markbates/goth"
)

type ConfigureController struct {
	beego.Controller
}

func (c *ConfigureController) ConfigureRepository() {
	cu := c.GetSession("current_user").(*goth.User)

	nwo := c.GetString("nwo")
	repository, err := github.GetRepository(cu.AccessToken, cu.NickName, nwo)
	if err != nil {
		beego.Error(err)
		c.Redirect("/404.html", 302)
		return
	}

	steps := newSteps("", "", "active", "disabled")
	steps.Select.Attr["owner"] = *repository.Owner.Login
	steps.Configure.Attr["nwo"] = nwo

	c.Data["steps"] = steps
	c.Data["currentUser"] = cu
	c.Data["repository"] = repository
	c.Data["templates"] = models.AllTemplates()

	c.TplName = "steps/configure.tpl"
}
