package models

import (
	"os"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func Init() error {
	orm.RegisterDataBase("default", "postgres", os.Getenv("DB_URL"), 30)

	orm.RegisterModel(new(Owner))
	orm.RegisterModel(new(Site))
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Template))

	if os.Getenv("DB_AUTO") != "" {
		if err := orm.RunSyncdb("default", true, true); err != nil {
			return err
		}
	}

	return loadTemplates()
}
