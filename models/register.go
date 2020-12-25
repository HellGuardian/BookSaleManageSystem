package models

import (
	"github.com/golang/glog"
)

type Admin struct {
	Name	string	`xorm:name`
	Tel		string	`xorm:tel`
	Sex		int		`xorm:sex`
	Age		int		`xorm:"age"`
	Email	string	`xorm:email`
	Role	int		`xorm:"role"`
	Password	string	`xorm:"password"`
}

func (a *Admin) RegisterUser(admin *Admin) error {
	// 事务处理
	session := engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		glog.Error("Begin register user session error,", err.Error())
		return err
	}

	// 插入数据库记录
	_, err := session.Insert(*admin)
	if err != nil {
		glog.Error("Register user error,", err.Error())
		session.Rollback()
		return err
	}
	if err := session.Commit(); err != nil {
		glog.Error("commit register info error,", err.Error())
		session.Rollback()
		return err
	}
	glog.Info("Register User Success.")
	return err
}
