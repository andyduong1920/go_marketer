package routers

import (
	"goMarketer/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/register", &controllers.RegistrationController{})
}
