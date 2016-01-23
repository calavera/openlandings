package models

type Site struct {
	ID          int64
	Name        string
	Description string
	Template    string
	Version     string
	Domain      string
	Analytics   string
	User        *User  `orm:"rel(fk)"`
	Owner       *Owner `orm:"rel(fk)"`
}
