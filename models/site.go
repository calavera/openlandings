package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/calavera/openlandings/github"
)

type Site struct {
	ID          int64
	Title       string
	Description string
	GitHubURL   string
	Template    string
	Version     string
	Domain      string
	Analytics   string
	User        *User     `orm:"rel(fk)"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
}

func CreateSite(user User, repository *github.Repository, template, domain string) error {
	site := Site{
		Title:     *repository.FullName,
		GitHubURL: *repository.HTMLURL,
		Template:  template,
		Domain:    domain,
		User:      &user,
	}

	if repository.Description != nil {
		site.Description = *repository.Description
	}

	o := orm.NewOrm()
	_, err := o.Insert(&site)
	return err
}
