package routers

import (
	"github.com/astaxie/beego"
	"github.com/calavera/openlandings/controllers"
	"github.com/calavera/openlandings/filters"
	"github.com/calavera/openlandings/models"
)

func init() {
	filters.Init()
	models.Init()

	beego.Router("/", &controllers.HomeController{})

	beego.Router("/login", &controllers.LoginController{}, "get:NewLogin")
	beego.Router("/auth/callback", &controllers.LoginController{}, "get:Callback")
}
