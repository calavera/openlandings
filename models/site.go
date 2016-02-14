package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/calavera/openlandings/github"
	githubapi "github.com/google/go-github/github"
)

type Site struct {
	ID          int64
	Title       string
	Description string
	GitHubURL   string `orm:"unique"`
	Template    string
	Version     string
	Domain      string `orm:"unique"`
	Analytics   string
	User        *User     `orm:"rel(fk)"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
}

type CombinedSite struct {
	Repository githubapi.Repository
	Site       *Site
}

type Sites struct {
	*github.Repositories
	All []*CombinedSite
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

func SiteByOwner(siteID, userID int64) (*Site, error) {
	var site Site
	o := orm.NewOrm()

	err := o.QueryTable("site").Filter("ID", siteID).Filter("User__ID", userID).One(&site)
	return &site, err
}

func LoadSites(repos *github.Repositories) (*Sites, error) {
	var sites []*Site
	var urls []interface{}
	for _, r := range repos.List {
		urls = append(urls, *r.HTMLURL)
	}

	o := orm.NewOrm()
	_, err := o.QueryTable("site").Filter("GitHubURL__in", urls...).All(&sites)
	if err != nil {
		return nil, err
	}

	configured := make(map[string]*Site)
	for _, s := range sites {
		configured[s.GitHubURL] = s
	}

	var combined []*CombinedSite
	for _, r := range repos.List {
		var site *Site
		if s, ok := configured[*r.HTMLURL]; ok {
			site = s
		}
		s := &CombinedSite{
			Repository: r,
			Site:       site,
		}
		combined = append(combined, s)
	}

	return &Sites{
		Repositories: repos,
		All:          combined,
	}, nil
}

func (s *Site) Delete() error {
	o := orm.NewOrm()
	_, err := o.Delete(s)
	return err
}
