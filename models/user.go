package models

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/markbates/goth"

	//	"github.com/markbates/goth"
)

type User struct {
	ID        int64
	Name      string
	AvatarURL string
	Slug      string    `orm:"unique"`
	UUID      string    `orm:"unique"`
	Email     string    `orm:"unique"`
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
}

func GetUser(slug string) (User, error) {
	o := orm.NewOrm()
	u := User{Slug: slug}
	err := o.Read(&u, "slug")
	return u, err
}

func RegisterUser(user *goth.User) error {
	o := orm.NewOrm()

	u := User{
		Name:      user.Name,
		Slug:      user.NickName,
		AvatarURL: user.AvatarURL,
		Email:     user.Email,
		UUID:      timeOrderedUUID(),
	}

	_, _, err := o.ReadOrCreate(&u, "Slug")
	return err
}

func timeOrderedUUID() string {
	unix := uint32(time.Now().UTC().Unix())

	b := make([]byte, 12)
	n, err := rand.Read(b)
	if n != len(b) {
		err = fmt.Errorf("Not enough entropy available")
	}
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%04x%08x",
		unix, b[0:2], b[2:4], b[4:6], b[6:8], b[8:])
}
