package controllers

import (
	"github.com/astaxie/beego"
	"github.com/calavera/openlandings/github"
	githubapi "github.com/google/go-github/github"
	"github.com/markbates/goth"
)

type StepsController struct {
	beego.Controller
}

func (c *StepsController) BrowseRepositories() {
	cu := c.GetSession("current_user").(*goth.User)
	u := c.Ctx.Input.GetData("github_user").(*githubapi.User)
	orgs, err := github.ListOrganizations(cu.AccessToken)
	if err != nil {
		beego.Error(err)
		c.Redirect("/404.html", 302)
		return
	}

	c.Data["steps"] = steps{
		Browse:    newStep("current"),
		Select:    newStep("disabled"),
		Configure: newStep("disabled"),
		Publish:   newStep("disabled"),
	}

	c.Data["currentUser"] = cu
	c.Data["githubUser"] = u
	c.Data["organizations"] = orgs

	c.TplName = "steps/browse.tpl"
}

type step struct {
	Status string
}

func newStep(status string) step {
	return step{status}
}

type steps struct {
	Browse    step
	Select    step
	Configure step
	Publish   step
}
