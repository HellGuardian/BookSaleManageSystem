package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "admin/login.tpl"
}

func (c *LoginController) Post() {
	userName := c.GetString("username")
	passWord := c.GetString("password")

	fmt.Println("username: ", userName)
	fmt.Println("password: ", passWord)
}
