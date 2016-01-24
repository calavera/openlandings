package models

import "time"

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
	Owner       *Owner    `orm:"rel(fk)"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
}
