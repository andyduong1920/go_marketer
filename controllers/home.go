package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	beego.ReadFromRequest(&c.Controller)

	c.Layout = "layouts/main.tpl"
	c.TplName = "home/index.tpl"

	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Notification"] = "layouts/notification.tpl"
}
