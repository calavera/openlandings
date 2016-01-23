package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func filterUser(ctx *context.Context) {
	u := ctx.Input.Session("current_user")
	if u == nil {
		ctx.Redirect(302, "/login")
	}
}

func Init() {
	beego.InsertFilter("/sites/*", beego.BeforeRouter, filterUser)
}
