package routers

import (
	"news/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {

	beego.InsertFilter("/Article/*",beego.BeforeRouter,filtFunc)
	beego.Router("/", &controllers.LoginController{},"get:Get;Post:HandleLogin")
	//退出
	beego.Router("/Article/Logout",&controllers.LoginController{},"*:Logout")
	//文章操作
	beego.Router("/Article/ShowArticle", &controllers.ArticleController{},"*:ShowArticleList")
	beego.Router("/Article/AddArticle", &controllers.ArticleController{},"get:ShowAddArticle;post:HandlerAddArticle")
	beego.Router("/Article/UpdateArticle", &controllers.ArticleController{},"get:ShowUpdate;post:UpdateArticle")
	beego.Router("/Article/ArticleContent", &controllers.ArticleController{},"get:ShowContent")
	beego.Router("/Article/DeleteArticle", &controllers.ArticleController{},"get:HandleDelete")
	beego.Router("/Article/UpdateArticle", &controllers.ArticleController{},"get:UpdateArticle")

    //注册
	beego.Router("/register", &controllers.RegisterController{},"get:ShowGet")
	beego.Router("/register", &controllers.RegisterController{},"post:HanleReg")
	//文章类型
	beego.Router("/Article/ShowTypeList",&controllers.TypeController{},"get:ShowTypeList;post:ShowType")
    beego.Router("/Article/AddType",&controllers.TypeController{},"get:ShowType;post:AddType")
	beego.Router("/Article/DeleteType", &controllers.TypeController{},"get:HandleDelete")
}

var filtFunc= func(ctx *context.Context) {
	userName:=ctx.Input.Session("userName")
	if userName ==nil {
		ctx.Redirect(302,"/")  //如果有输出不在执行
	}
}