package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/calavera/openlandings/github"
	"github.com/markbates/goth"
)

func filterUser(ctx *context.Context) {
	u := ctx.Input.Session("current_user")
	if u == nil {
		ctx.Redirect(302, "/login")
	}

	if gu := ctx.Input.GetData("github_user"); gu == nil {
		g, err := github.GetCurrentUser(u.(*goth.User).AccessToken)
		if err != nil {
			ctx.Redirect(302, "/404.html")
		}
		ctx.Input.SetData("github_user", g)
	}
}

func Init() {
	beego.InsertFilter("/steps/*", beego.BeforeRouter, filterUser)
}
