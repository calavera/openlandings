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
	filePath := c.GetString("load")
	if filePath == "" {
		filePath = "README.md"
	}
	repository, err := github.GetRepositoryWithContent(cu.AccessToken, cu.NickName, nwo, filePath)
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
