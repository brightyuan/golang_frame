package main

import (
	"openapi/controllers"
	_ "openapi/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Router("/v1/user", &controllers.UserController{})
	beego.Run()
}
