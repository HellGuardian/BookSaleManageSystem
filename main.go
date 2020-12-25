package main

import (
	"BookSaleManageSystem/models"
	_ "BookSaleManageSystem/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.InitData()
	beego.Run()
}

