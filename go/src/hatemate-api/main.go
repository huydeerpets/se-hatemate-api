package main

import (
	_ "fmt"
	_ "hatemate-api/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
