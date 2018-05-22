package routers

import (
	"hatemate-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/GetUser", &controllers.UserController{}, "*:Get")
	beego.Router("/RegisterUser", &controllers.UserController{}, "*:Insert")
	beego.Router("/GetAllUsers", &controllers.UserController{}, "*:GetAllUserController")
	beego.Router("/AddNewTopic", &controllers.UserController{}, "*:AddNewTopic")
	beego.Router("/GetAllTopics", &controllers.UserController{}, "*:GetAllTopicController")
}
