package controllers

import(
	"github.com/astaxie/beego"
)

type WelComeController struct {
	beego.Controller
}

func (c *WelComeController) Get() {
	c.TplName = "admin/welcome.html"
}
