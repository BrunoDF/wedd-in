package main

import (
	"wedd-in/models"
	_ "wedd-in/routers"

	"github.com/astaxie/beego"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	// Swagger
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// Database startup
	models.InitDB()

	// Database close on end
	defer models.DB.Close()

	// Run
	beego.Run()

}
