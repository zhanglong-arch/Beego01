package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "https://www.baidu.com"
	c.Data["Email"] = "1621258931@qq.com"
	c.TplName = "index.tpl"
}
