package models

import (
	"encoding/json"
	"path/filepath"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/calavera/openlandings/themes"
)

type Template struct {
	ID          int64  `json:"-"`
	Path        string `json:"-",orm:"unique"`
	BasePath    string `json:"-",orm:"unique"`
	Name        string `json:"name",orm:"unique"`
	Home        string `json:"home"`
	Description string `json:"description"`
}

func AllTemplates() []*Template {
	var templates []*Template

	o := orm.NewOrm()
	_, err := o.QueryTable("template").All(&templates)
	if err != nil {
		beego.Error(err)
	}

	return templates
}

func GetTemplate(basePath string) (Template, error) {
	o := orm.NewOrm()
	t := Template{BasePath: basePath}
	err := o.Read(&t, "base_path")
	return t, err
}

func loadTemplates() error {
	th, err := themes.AllFromDisk()
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	for _, t := range th {
		var template Template
		j, err := t.JSONFile()
		if err != nil {
			return err
		}
		if err := json.NewDecoder(j).Decode(&template); err != nil {
			return err
		}

		template.Path = t.Path
		template.BasePath = filepath.Base(t.Path)
		if _, _, err := o.ReadOrCreate(&template, "name"); err != nil {
			return err
		}
	}

	return nil
}
