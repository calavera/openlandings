package filters

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/calavera/openlandings/github"
	"github.com/calavera/openlandings/models"
	"github.com/markbates/goth"
)

func filterUser(ctx *context.Context) {
	u, ok := ctx.Input.Session("current_user").(*goth.User)
	if !ok || u == nil {
		ctx.Redirect(302, "/login")
	}

	if gu := ctx.Input.GetData("github_user"); gu == nil {
		g, err := github.GetCurrentUser(u.AccessToken)
		if err != nil {
			beego.Error(err)
			ctx.Redirect(302, "/404.html")
		}
		ctx.Input.SetData("github_user", g)
	}
}

func filterSiteOwner(ctx *context.Context) {
	u := ctx.Input.Session("current_user").(*goth.User)

	user, err := models.GetUser(u.NickName)
	if err != nil {
		beego.Error(err)
		ctx.Redirect(302, "/404.html")
	}

	siteID, err := strconv.ParseInt(ctx.Input.Param(":id"), 10, 64)
	if err != nil {
		beego.Error(err)
		ctx.Redirect(302, "/404.html")
	}

	site, err := models.SiteByOwner(siteID, user.ID)
	if err != nil {
		beego.Error(err)
		ctx.Redirect(302, "/404.html")
	}
	ctx.Input.SetData("user", user)
	ctx.Input.SetData("site", site)
}

func filterPostMethods(ctx *context.Context) {
	if ctx.Input.Query("_method") != "" && ctx.Input.IsPost() {
		ctx.Request.Method = ctx.Input.Query("_method")
	}
}

func Init() {
	beego.InsertFilter("*", beego.BeforeRouter, filterPostMethods)

	beego.InsertFilter("/steps/*", beego.BeforeExec, filterUser)
	beego.InsertFilter("/sites/*", beego.BeforeExec, filterUser)
	beego.InsertFilter("/sites/*", beego.BeforeExec, filterSiteOwner)
}
