package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
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

	password := params["password"]
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.HashedPassword = string([]byte(hashedPassword))

	valid := validate(user, password)

	if valid.HasErrors() {
		return valid.Errors[0]
	} else {
		_, err := o.Insert(user)

		return err
	}
}

func validate(user *User, password string) validation.Validation {
	validator := validation.Validation{}

	validator.Email(user.Email, "email")

	validator.MinSize(password, 6, "password")
	validator.MaxSize(password, 10, "password")

	return validator
}
