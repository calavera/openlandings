package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/astaxie/beego"
	githubapi "github.com/calavera/openlandings/github"
	"github.com/calavera/openlandings/models"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
)

var provider goth.Provider

func init() {
	url := os.Getenv("SITE_URL") + "/auth/callback"
	provider = github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), url, "user:read", "user:email", "read:org")
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) NewLogin() {
	c.getAuthToken(c.Ctx.ResponseWriter, c.Ctx.Request)
}

func (c *LoginController) Callback() {
	user, err := c.completeAuth()
	if err != nil {
		beego.Error(err)
	} else {
		if err := models.RegisterUser(user); err != nil {
			beego.Error(err)
			http.Error(c.Ctx.ResponseWriter, err.Error(), 500)
			return
		}
		c.SetSession("current_user", user)
	}
	http.Redirect(c.Ctx.ResponseWriter, c.Ctx.Request, "/", http.StatusTemporaryRedirect)
}

func (c *LoginController) getAuthToken(w http.ResponseWriter, req *http.Request) error {
	sess, err := provider.BeginAuth(gothic.SetState(req))
	if err != nil {
		return err
	}

	url, err := sess.GetAuthURL()
	if err != nil {
		return err
	}
	c.SetSession("openl_login_session", sess.Marshal())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil
	}

	http.Redirect(w, req, url, http.StatusTemporaryRedirect)
	return nil
}

func (c *LoginController) completeAuth() (*goth.User, error) {
	s := c.GetSession("openl_login_session")
	if s == nil {
		http.Redirect(c.Ctx.ResponseWriter, c.Ctx.Request, "/login", http.StatusTemporaryRedirect)
		return nil, fmt.Errorf("unknown session")
	}
	sess, err := provider.UnmarshalSession(s.(string))
	if err != nil {
		return nil, err
	}

	_, err = sess.Authorize(provider, c.Ctx.Request.URL.Query())
	user, err := provider.FetchUser(sess)
	if err != nil {
		return nil, err
	}
	user.NickName = user.RawData["login"].(string)

	if user.Email == "" {
		user.Email = githubapi.GetPrimaryEmail(user.AccessToken)
	}
	return &user, nil
}
