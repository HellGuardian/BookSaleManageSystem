package controllers

import (
	"BookSaleManageSystem/models"
	"github.com/astaxie/beego"
	"github.com/golang/glog"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Get() {
	r.TplName = "web/register.html"
}

func (r *RegisterController) Post() {
	admin := &models.Admin{}
	admin.Name = r.GetString("username")
	admin.Password = r.GetString("repass")
	admin.Tel = r.GetString("phone")
	sex := r.GetString("sex")
	if sex == "男" {
		admin.Sex = 1
	} else {
		admin.Sex = 0
	}
	admin.Email = r.GetString("email")
	age, err := r.GetInt("age")
	if err != nil {
		glog.Error("get user age error,", err.Error())
		return
	}
	admin.Age = age
	admin.Role = 2

	if err := admin.RegisterUser(admin); err != nil {
		glog.Error("call register func error,", err.Error())
		return
	}
}
