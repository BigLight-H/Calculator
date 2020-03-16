package main

import (
	_ "Calculator/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/static", "static")

	beego.Run()
}

