package main

import (
	_ "hatemate-api/models"
	_ "hatemate-api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase(
		"default",
		"postgres",
		"user=HateMate password=HateMateApp123 host=35.198.207.177 port=5432 dbname=HateMateDB sslmode=disable")

	orm.RunSyncdb("default", false, true)
}

func main() {
	orm.Debug = true

	o := orm.NewOrm()
	o.Using("default")

	// TestUser = models.hatemateusers.HateMateUser
	// TestUser.ID = 0
	// TestUser.Email = "dummy@dummy.com"
	// TestUser.Password = "dummy"
	// TestUser.FirstName = "Fon"
	// TestUser.LastName = "Magenus"
	// TestUser.Gender = "Male"

	// id, err := o.Insert(TestUser)
	// fmt.Printf("ID: %d, ERR: %v\n", id, err)
	beego.Run()
}
