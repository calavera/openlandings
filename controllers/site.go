package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/calavera/openlandings/hosting"
	"github.com/calavera/openlandings/models"
)

type SiteController struct {
	beego.Controller
}

func (c *SiteController) Delete() {
	site := c.Ctx.Input.GetData("site").(*models.Site)
	user := c.Ctx.Input.GetData("user").(models.User)

	if err := hosting.Delete(site, user); err != nil {
		beego.Error(err)
		c.Redirect("/404.html", 302)
		return
	}

	if err := site.Delete(); err != nil {
		beego.Error(err)
		c.Redirect("/404.html", 302)
		return
	}

	selectURL := fmt.Sprintf("/steps/select?owner=%s", user.Slug)
	c.Redirect(selectURL, 302)
}
