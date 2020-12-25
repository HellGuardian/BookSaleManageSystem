package routers

import (
	"BookSaleManageSystem/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.IndexController{})
	beego.Router("/welcome.html", &controllers.WelComeController{})
	beego.Router("/login.html", &controllers.LoginController{})
    beego.Router("/order-list.html", &controllers.OrderController{})
	beego.Router("/book-list.html", &controllers.BooksController{})
    beego.Router("/admin-list.html", &controllers.AdminListController{})

	// 前台
	beego.Router("/index.html", &controllers.WebController{})
	beego.Router("/register.html", &controllers.RegisterController{})
}
