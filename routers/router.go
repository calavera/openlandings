package routers

import (
	"github.com/astaxie/beego"
	"github.com/calavera/openlandings/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})

	beego.Router("/login", &controllers.LoginController{}, "get:NewLogin")
	beego.Router("/auth/callback", &controllers.LoginController{}, "get:Callback")
}
