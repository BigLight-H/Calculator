package routers

import (
	"Calculator/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.IndexController{},"get:Get")
}
