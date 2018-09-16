package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"news/models"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) ShowGet(){
this.TplName="register.html"
}
func (this *RegisterController) HanleReg()  {
	name:=this.GetString("userName")
	passwd:=this.GetString("password")
	if name=="" || passwd =="" {
		beego.Info("用户名密码不能为空")
		this.TplName="register.html"
	}

	user:=new(models.User)
	user.Passwd=passwd
	user.UserName=name
	o:= orm.NewOrm()
	_,err:=o.Insert(user)
	if err!=nil {
		beego.Info("Registererr",err)
		return
	}
	beego.Info(name,passwd)
	beego.Info("注册成功")
	//this.TplName=
	this.Redirect("/",302)
}


