package controllers

import (
	"github.com/astaxie/beego"
)

type BooksController struct {
	beego.Controller
}

func (c *BooksController) Get() {
	c.TplName = "admin/book-list.html"
}
