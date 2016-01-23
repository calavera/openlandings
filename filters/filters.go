package filters

import (
	"github.com/astaxie/beego"
	"golang.org/x/net/context"
)

func filterUser(ctx *context.Context) {
	_, ok := ctx.Input.Session("current_user")
	if !ok {
		ctx.Redirect(302, "/login")
	}
}

func Init() {
	beego.InsertFilter("/sites/*", beego.BeforeRouter, filterUser)
}
