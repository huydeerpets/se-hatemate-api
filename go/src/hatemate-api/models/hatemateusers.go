package models

import (
	"github.com/astaxie/beego/orm"
)

type HateMateUser struct {
	ID        int    `json:"user_id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
}

func init() {
	orm.RegisterModel(new(HateMateUser))
}

func GetAllUser() []HateMateUser {
	o := orm.NewOrm()
	var users []HateMateUser
	_, _ = o.QueryTable("hate_mate_user").All(&users)
	return users
}
