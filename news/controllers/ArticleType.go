package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"news/models"
)

type TypeController struct {
	beego.Controller
}

func(this *TypeController) ShowType()  {
	this.TplName="addType.html"
}

func (this *TypeController)ShowTypeList()  {
/*	session:=this.StartSession()
	beego.Info("/ShowTypeList GetSession:")
	username:=session.Get("username")
	beego.Info("username:",username)
	if username=="" {
		beego.Info("无效登录！")
		this.TplName="addType.html"
		return
	}*/

	o:=orm.NewOrm()
	qs:=o.QueryTable("ArticleType")
	var types []models.ArticleType

	qs.All(&types)
	this.Data["types"]=types
	this.Layout="layout.html"
	this.TplName="addType.html"
}

func (this *TypeController) AddType()  {
	typeName:=this.GetString("typeName")
	if typeName=="" {
		beego.Info("类型名称不能为空！")
		this.TplName="addType.html"
		return
	}
	o:=orm.NewOrm()
	types:=models.ArticleType{}
	types.TypeName=typeName

	_,err:=o.Insert(&types)
	if err!=nil {
		beego.Info("TypeController Insert err:",err)
		this.TplName="addType.html"
		return
	}
	this.Redirect("/Article/ShowTypeList",302)
}

func (this *TypeController) HandleDelete()  {
	id,_:=this.GetInt("id")
	o:=orm.NewOrm()
	types:=models.ArticleType{}
	types.Id = id
	_,err:=o.Delete(&types)
	if err!=nil {
		beego.Info("Delete err:",err)
		this.Redirect("/Article /ShowTypeList",302)
		return
	}
	this.Redirect("/Article/ShowTypeList",302)

}