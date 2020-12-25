package controllers

import (
	"github.com/astaxie/beego"
)

type WebController struct {
	beego.Controller
}

func (c *WebController) Get() {
	c.Data["Pic_Path"] = "./static/book_pic/2.pic.jpg"
	c.TplName = "web/index.tpl"
}
