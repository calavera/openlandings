package controllers

import (
	"fmt"
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/calavera/openlandings/github"
	"github.com/calavera/openlandings/hosting"
	"github.com/calavera/openlandings/models"
	"github.com/calavera/openlandings/themes"
	"github.com/markbates/goth"
)

type PublishController struct {
	beego.Controller
}

func (c *PublishController) ConfigureSite() {
	cu := c.GetSession("current_user").(*goth.User)

	f := PublishForm{}
	if err := c.ParseForm(&f); err != nil {
		beego.Error(err)
		c.Redirect("/404.html", 302)
		return
	}

	repository, err := github.GetRepository(cu.AccessToken, cu.NickName, f.Nwo)
	if err != nil {
		beego.Error(err)
		c.Redirect("/404.html", 302)
		return
	}

	steps := newSteps("", "", "", "active")
	steps.Select.Attr["owner"] = *repository.Owner.Login
	steps.Configure.Attr["nwo"] = f.Nwo

	c.Data["steps"] = steps
	c.Data["currentUser"] = cu
	c.Data["preview"] = f
	c.Data["repository"] = repository

	c.TplName = "steps/publish.tpl"
}

func (c *PublishController) PublishSite() {
	cu := c.GetSession("current_user").(*goth.User)

	user, err := models.GetUser(cu.NickName)
	if err != nil {
		beego.Error(err)
		c.Redirect("/404.html", 302)
		return
	}

	f := PublishForm{}
	if err := c.ParseForm(&f); err != nil {
		beego.Error(err)
		c.Redirect("/404.html", 302)
		return
	}

	repository, err := github.GetRepository(cu.AccessToken, cu.NickName, f.Nwo)
	if err != nil {
		beego.Error(err)
		c.Redirect("/404.html", 302)
		return
	}

	tmpl, err := models.GetTemplate(f.Template)
	if err != nil {
		beego.Error(err)
		c.Redirect("/404.html", 302)
		return
	}

	siteName := siteName(repository)
	domain, fullURL := finalDomain(siteName, f)

	if err := models.CreateSite(user, repository, f.Template, fullURL); err != nil {
		beego.Error(err)
		c.Redirect("/404.html", 302)
		return
	}

	deployPath, err := themes.Pack(repository, tmpl.Path, fullURL, f.Landing)
	if err != nil {
		beego.Error(err)
		c.Redirect("/404.html", 302)
		return
	}
	defer os.RemoveAll(deployPath)

	if err := hosting.Publish(user, deployPath, siteName, domain); err != nil {
		beego.Error(err)
		c.Redirect("/404.html", 302)
		return
	}

	steps := newSteps("", "", "", "disabled")
	steps.Select.Attr["owner"] = *repository.Owner.Login
	steps.Configure.Attr["nwo"] = f.Nwo
	c.Data["steps"] = steps

	c.Data["currentUser"] = cu
	c.Data["preview"] = f
	c.Data["repository"] = repository
	c.Data["domainURL"] = fullURL

	c.TplName = "steps/finish.tpl"
}

type PublishForm struct {
	Landing      string `form:"landing"`
	Nwo          string `form:"nwo"`
	Template     string `form:"template"`
	CustomDomain string `form:"custom"`
}

func finalDomain(siteName string, f PublishForm) (string, string) {
	if f.CustomDomain != "" {
		return f.CustomDomain, fmt.Sprintf("https://%s", f.CustomDomain)
	}
	domain := fmt.Sprintf("%s.netlify.com", siteName)
	// normal domains must have this field empty.
	return "", fmt.Sprintf("http://%s", domain)
}

func siteName(repository *github.Repository) string {
	return strings.Replace(fmt.Sprintf("%s-%s", *repository.Owner.Login, *repository.Name), "_", "-", -1)
}
