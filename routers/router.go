package routers

import (
	"Beego01/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/beego", &controllers.MainController{})
}