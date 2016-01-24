package models

import (
	"encoding/json"

	"github.com/astaxie/beego/orm"
	"github.com/calavera/openlandings/themes"
)

type Template struct {
	ID          int64  `json:"-"`
	Path        string `json:"-"`
	Name        string `json:"name"`
	Home        string `json:"home"`
	Description string `json:"description"`
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
		if _, _, err := o.ReadOrCreate(&template, "name"); err != nil {
			return err
		}
	}

	return nil
}
