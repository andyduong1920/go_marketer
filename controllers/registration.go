package controllers

import (
	"fmt"
	"html/template"

	"goMarketer/models"

	"github.com/astaxie/beego"
)

type RegistrationController struct {
	beego.Controller
}

func (c *RegistrationController) Get() {
	beego.ReadFromRequest(&c.Controller)

	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())

	c.Layout = "layouts/auth.tpl"
	c.TplName = "registration/new.tpl"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Notification"] = "layouts/notification.tpl"
}

func (c *RegistrationController) Post() {
	flash := beego.NewFlash()

	user := new(models.User)
	user.Email = c.GetString("email")

	err := models.CreateUser(user, c.GetString("password"))

	if err != nil {
		// TODO: Re-render the form instead of redirect
		flash.Error(fmt.Sprint(err))
	} else {
		flash.Notice("User created successfully.")
	}

	flash.Store(&c.Controller)

	c.Ctx.Redirect(302, "/register")
}
