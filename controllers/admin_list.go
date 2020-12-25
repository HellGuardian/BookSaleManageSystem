package controllers

import (
	"github.com/astaxie/beego"
)

type AdminListController struct {
	beego.Controller
}

func (c *AdminListController) Get() {
	c.TplName = "admin/admin-list.html"
}