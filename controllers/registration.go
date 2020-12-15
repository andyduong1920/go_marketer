package controllers

import (
	"fmt"
	"goMarketer/models"
	_ "goMarketer/models"

	"github.com/astaxie/beego"
)

type RegistrationController struct {
	beego.Controller
}

func (c *RegistrationController) Get() {
	beego.ReadFromRequest(&c.Controller)

	c.Layout = "layouts/auth.tpl"
	c.TplName = "registration/new.tpl"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Notification"] = "layouts/notification.tpl"
}

func (c *RegistrationController) Post() {
	flash := beego.NewFlash()

	params := make(map[string]string)
	params["email"] = c.GetString("email")
	params["password"] = c.GetString("password")

	err := models.CreateUser(params)

	if err != nil {
		// TODO: Re-render the form instead of redirect
		flash.Error(fmt.Sprint(err))
	} else {
		flash.Notice("User created successfully.")
	}

	flash.Store(&c.Controller)

	c.Ctx.Redirect(302, "/register")
}
