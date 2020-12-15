package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	Email    string `orm:"unique"`
	Password string

	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "users"
}

func CreateUser(params map[string]string) (int64, error) {
	o := orm.NewOrm()

	user := new(User)
	user.Email = params["email"]
	user.Password = params["password"]

	id, err := o.Insert(user)

	return id, err
}
