package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"news/models"
	"time"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get()  {

	name:=this.Ctx.GetCookie("userName")

	this.Data["name"]=name
	this.TplName="login.html"
}

func (this *LoginController)HandleLogin()  {

	name:=this.GetString("userName")
	passwd:=this.GetString("password")
	if name=="" || passwd =="" {
		beego.Info("登录用户名密码不能为空")
		this.TplName="login.html"
	}
	user:=new(models.User)
	user.UserName=name
	o:=orm.NewOrm()
	err:=o.Read(user,"UserName")
	if err !=nil {
		beego.Info("Login UserName err:",err)
		this.TplName="login.html"
		return
	}
	if user.Passwd !=passwd {
		beego.Info("Login passwd err:",err)
		this.TplName="login.html"
		return
	}
	check:=this.GetString("remember")
	if check =="on" {
		this.Ctx.SetCookie("userName",name,time.Second*3600)
	}else {
		this.Ctx.SetCookie("userName",name,-1)
	}
	this.SetSession("userName",name)

	this.Redirect("/Article/ShowArticle",302)
}

func (this *LoginController)Logout()  {
	//this.Ctx.Redirect(302,"login.html")
	this.DelSession("userName")
	this.TplName="/"
}


