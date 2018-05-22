package controllers

import (
	"encoding/json"
	"hatemate-api/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// UserController test
type UserController struct {
	beego.Controller
}

// Get user
func (uc *UserController) Get() {
	var user models.HateMateUser
	json.Unmarshal(uc.Ctx.Input.RequestBody, &user)

	o := orm.NewOrm()
	o.Using("default")

	err := o.Read(&user, "Email", "Password")

	if err == orm.ErrNoRows {
		uc.Ctx.Output.SetStatus(404)
		uc.Ctx.Output.Body([]byte("User not found"))
		return
	}

	uc.Data["json"] = user
	uc.ServeJSON()
}

// Get All user
func (uc *UserController) GetAllUserController() {
	users := models.GetAllUser()
	uc.Data["json"] = users
	uc.ServeJSON()
}

// Insert user
func (uc *UserController) Insert() {
	var user models.HateMateUser
	json.Unmarshal(uc.Ctx.Input.RequestBody, &user)

	o := orm.NewOrm()
	o.Using("default")

	err := o.Read(&user, "email")

	if err == nil {
		uc.Ctx.Output.SetStatus(400)
		uc.Ctx.Output.Body([]byte("User already exists"))
		return
	}

	id, err := o.Insert(&user)

	if err != nil {
		uc.Ctx.Output.SetStatus(400)
		uc.Ctx.Output.Body([]byte("Error inserting user"))
		return
	}

	uc.Data["json"] = id
	uc.ServeJSON()
}
