package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/calavera/openlandings/github"
	"github.com/calavera/openlandings/models"
	githubapi "github.com/google/go-github/github"
	"github.com/markbates/goth"
)

type StepsController struct {
	beego.Controller
}

func (c *StepsController) BrowseOrganizations() {
	cu := c.GetSession("current_user").(*goth.User)
	u := c.Ctx.Input.GetData("github_user").(*githubapi.User)
	orgs, err := github.ListOrganizations(cu.AccessToken)
	if err != nil {
		beego.Error(err)
		c.Redirect("/404.html", 302)
		return
	}

	c.Data["steps"] = newSteps("active", "disabled", "disabled", "disabled")
	c.Data["currentUser"] = cu
	c.Data["githubUser"] = u
	c.Data["organizations"] = orgs

	c.TplName = "steps/browse.tpl"
}

func (c *StepsController) BrowseRepositories() {
	cu := c.GetSession("current_user").(*goth.User)
	u := c.Ctx.Input.GetData("github_user").(*githubapi.User)

	var ownerWithRepos *github.Repositories
	var err error

	owner := c.GetString("owner")
	page := c.getPage()
	if owner == "" || owner == *u.Login {
		ownerWithRepos, err = getUserRepositories(page, cu)
	} else {
		ownerWithRepos, err = getOrgRepositories(page, owner, cu)
	}

	if err != nil {
		beego.Error(err)
		c.Redirect("/404.html", 302)
		return
	}

	sites, err := models.LoadSites(ownerWithRepos)
	if err != nil {
		beego.Error(err)
		c.Redirect("/404.html", 302)
		return
	}

	steps := newSteps("", "active", "disabled", "disabled")
	steps.Select.Attr["owner"] = owner

	c.Data["steps"] = steps
	c.Data["currentUser"] = cu
	c.Data["sites"] = sites
	c.Data["owner"] = u
	if ownerWithRepos.Owner != nil {
		c.Data["owner"] = ownerWithRepos.Owner
	}

	c.TplName = "steps/select.tpl"
}

func getUserRepositories(page int, u *goth.User) (*github.Repositories, error) {
	opts := &githubapi.RepositoryListOptions{
		Type: "owner",
		ListOptions: githubapi.ListOptions{
			Page: page,
		},
	}
	return github.ListUserRepositories(u.AccessToken, opts)
}

func getOrgRepositories(page int, owner string, u *goth.User) (*github.Repositories, error) {
	opts := &githubapi.RepositoryListByOrgOptions{
		Type: "public",
		ListOptions: githubapi.ListOptions{
			Page: page,
		},
	}
	return github.ListOrgRepositories(u.AccessToken, owner, opts)
}

func (c *StepsController) getPage() int {
	o := c.GetString("page")
	if o == "" {
		return 1
	}
	page, err := strconv.Atoi(o)
	if err != nil {
		page = 1
	}
	return page
}

type step struct {
	Status string
	Attr   map[string]string
}

func newStep(status string) step {
	return step{
		Status: status,
		Attr:   make(map[string]string),
	}
}

type steps struct {
	Browse    step
	Select    step
	Configure step
	Publish   step
}

func newSteps(browse, selectStep, configure, publish string) steps {
	return steps{
		Browse:    newStep(browse),
		Select:    newStep(selectStep),
		Configure: newStep(configure),
		Publish:   newStep(publish),
	}
}
