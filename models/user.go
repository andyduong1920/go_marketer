package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id             int
	Email          string `orm:"unique"`
	HashedPassword string

	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) TableName() string {
	return "users"
}

func CreateUser(params map[string]string) error {
	o := orm.NewOrm()

	user := new(User)
	user.Email = params["email"]

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(params["password"]), bcrypt.DefaultCost)
	user.HashedPassword = string([]byte(hashedPassword))

	_, err := o.Insert(user)

	return err
}
