package models

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"ossDevOps/utils"
)

var engine *xorm.Engine

// 初始化参数和数据库
func InitData() (err error) {
	dbUrl := beego.AppConfig.String("dbUrl")
	if dbUrl == "" {
		glog.Error("Connect MySql url null.")
		return
	}
	engine, err = xorm.NewEngine("mysql", dbUrl)
	if err != nil {
		glog.Error("New MySql engine error", err)
		return
	}

	var maxIdleConn int
	maxIdleConn, err = beego.AppConfig.Int("maxIdleConns")
	if err != nil {
		glog.Error("Get config maxIdleConns error.", err)
		return err
	}
	engine.SetMaxIdleConns(maxIdleConn)

	var maxOpenConn int
	maxOpenConn, err = beego.AppConfig.Int("maxOpenConns")
	if err != nil {
		utils.FileLog.Error("Get config maxOpenConns error.", err)
		return err
	}
	engine.SetMaxOpenConns(maxOpenConn)

	err = engine.Ping()
	if err != nil {
		glog.Error("MySql ping error,", err)
		return err
	}

	return err
}

