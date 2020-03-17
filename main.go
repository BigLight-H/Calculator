package main

import (
	"Calculator/models"
	_ "Calculator/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.Debug = true
	models.Init()
	beego.SetStaticPath("/static", "static")

	beego.Run()
}

