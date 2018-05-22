package routers

import (
	"hatemate-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/GetUser", &controllers.UserController{}, "*:Get")
	beego.Router("/RegisterUser", &controllers.UserController{}, "*:Insert")
}
