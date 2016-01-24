package routers

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/calavera/openlandings/controllers"
	"github.com/calavera/openlandings/filters"
	"github.com/calavera/openlandings/models"
)

func init() {
	if err := models.Init(); err != nil {
		log.Fatal(err)
	}
	filters.Init()

	beego.Router("/", &controllers.HomeController{})
	beego.Router("/steps/browse", &controllers.StepsController{}, "get:BrowseOrganizations")
	beego.Router("/steps/select", &controllers.StepsController{}, "get:BrowseRepositories")
	beego.Router("/steps/configure", &controllers.ConfigureController{}, "get:ConfigureRepository")

	beego.Router("/login", &controllers.LoginController{}, "get:NewLogin")
	beego.Router("/auth/callback", &controllers.LoginController{}, "get:Callback")
}
