package controllers

import (
	"github.com/astaxie/beego"
)

// MainController test
type MainController struct {
	beego.Controller
}

// Get test
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
