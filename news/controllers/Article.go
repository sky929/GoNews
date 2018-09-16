package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"news/models"
	"path"
	"strconv"
	"time"
)

type ArticleController struct {
	beego.Controller
}

//处理下来框改变的值
func (this *ArticleController) HandleSelect() {
	typeName := this.GetString("select")
	if typeName == "" {
		return
	}
	o := orm.NewOrm()
	var artiles []models.ArticleType
	o.QueryTable("artiles").RelatedSel("artiles").Filter("NewsTpye__TypeName", typeName).All(&artiles)
}

func (this *ArticleController) ShowArticleList() {
	//
/*	userName := this.GetSession("userName")

	if userName == "" {
		this.Redirect("/", 302)
	}
*/
	pageIndex, _ := this.GetInt("pageIndex")
	if pageIndex == 0 {
		pageIndex = 1
	}
	o := orm.NewOrm()
	article := new(models.Article)
	qs := o.QueryTable(article)

	//文章类型绑定
	typeName := this.GetString("select")

	types := new([]models.ArticleType)
	o.QueryTable("ArticleType").All(types)
	this.Data["types"] = types
	//下拉列表绑定选定值
	this.Data["select"] = typeName

	var articles []models.Article
	pageSize := 4
	count, _ := qs.RelatedSel("ArticleType").Filter("ArticleType__TypeName",typeName).Count()
	pageCount := math.Ceil(float64(count) / float64(pageSize))

	if typeName == "" {
		qs.Limit(pageSize, pageSize*(pageIndex-1)).RelatedSel("ArticleType").All(&articles) //惰性查询
	} else {
		qs.Limit(pageSize, pageSize*(pageIndex-1)).RelatedSel("ArticleType").Filter("ArticleType__TypeName", typeName).All(&articles) //
		//o.QueryTable("artiles").RelatedSel("artiles").Filter("NewsTpye__TypeName",typeName).All(&artiles)
	}

	this.Data["articles"] = articles

	FristPage := false //标记首页
	EndPage := false   //标记尾页
	if pageIndex == 1 {
		FristPage = true
	}
	if pageIndex == int(pageCount) {
		EndPage = true
	}

	this.Data["count"] = count         //总行数
	this.Data["pageCount"] = pageCount //总页数pageCount

	this.Data["pageIndex"] = pageIndex // math.Ceil(float64(count) / float64(pageSize))
	this.Data["FristPage"] = FristPage
	this.Data["EndPage"] = EndPage




	userName := this.GetSession("userName")
	this.Data["userName"]=userName
	this.Layout = "layout.html"
	this.TplName = "index.html"
}

func (this *ArticleController) ShowAddArticle() {
	//文章类型绑定
	o := orm.NewOrm()
	types := new([]models.ArticleType)
	o.QueryTable("ArticleType").All(types)
	this.Data["types"] = types
	this.Layout="layout.html"
	this.TplName = "add.html"
}

func (this *ArticleController) HandlerAddArticle() {
	title := this.GetString("articleName")
	content := this.GetString("content")
	//上传图片
	f, h, err := this.GetFile("uploadname")
	if err != nil {
		beego.Info("GetFile err:", err)
		this.Redirect("Article/AddArticle", 302) //= "add.html"
		return
	}
	defer f.Close()
	//1.文件格式
	ext := path.Ext(h.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		beego.Info("path.Ext 上传格式不正确", err)
		this.TplName = "add.html"
		return
	}
	//2.文件大小
	if h.Size > 409600 {
		beego.Info("上传文件太大，不允许上传！ ", err)
		this.TplName = "add.html"
		return
	}
	//3不能重名F:\GoCode\src\news\static\img
	//filename:=time.Now().Format("2006-01-02 15:04:05")
	filename := strconv.Itoa(time.Now().Nanosecond())
	err = this.SaveToFile("uploadname", "./static/img/"+filename+ext)
	//beego.Info("/static/img/"+filename+ext)

	if err != nil {
		beego.Info("SaveToFile err:", err)
		this.TplName = "Article/add.html"
		return
	}

	o := orm.NewOrm()
	artice := models.Article{}
	artice.Title = title
	artice.Content = content
	artice.Img = "./static/img/" + filename + ext

	//获取文章类型
	typeName := this.GetString("select")
	if typeName == "" {
		beego.Info("下拉框未选择")
		return
	}
	var articleType models.ArticleType
	articleType.TypeName = typeName
	o.Read(&articleType, "TypeName")

	artice.ArticleType = &articleType
	beego.Info("articleType:", articleType.Id, "   ", articleType.TypeName)
	_, err = o.Insert(&artice)
	if err != nil {
		beego.Info("Insert err:", err)
		this.TplName = "add.html"
		return
	}
	//this.Ctx.WriteString("新闻发布成功")
	this.Redirect("ShowArticle", 302)
}

func (this *ArticleController) ShowContent() {
	id, _ := this.GetInt("id")
	o := orm.NewOrm()
	artilce := models.Article{}
	artilce.Id = id
	err := o.Read(&artilce)
	if err != nil {
		beego.Info("Read err:", err)
		this.Redirect("Article/ShowArticle", 302)
		return
	}
	artilce.Count += 1
	//多对多插入duzhe
	//artilcs:=models.Article{Id:id}
	m2m := o.QueryM2M(&artilce, "Users")
	userName := this.GetSession("userName")

	user := models.User{}
	user.UserName = userName.(string)
	o.Read(&user, "UserName")
	beego.Info("===========================DEBUG================================")
	beego.Info("userName:",user)

	_, err = m2m.Add(&user)

	if err != nil {
		beego.Info("m2m.Add :", err)
	}
	o.Update(&artilce)

	//多对多查询
	users := []models.User{}
	//o.LoadRelated(&artilc,"Users")
	count,_:=o.QueryTable("User").Filter("Articles__Article__Id", id).Distinct().All(&users)

	beego.Info("===========================DEBUG================================")
	beego.Info("count:",count)
	beego.Info("users:",users)
	this.Data["users"] = users
	this.Data["artilce"] = artilce
	this.Layout="layout.html"
	this.TplName = "content.html"
}

func (this *ArticleController) HandleDelete() {
	id, _ := this.GetInt("id")
	o := orm.NewOrm()
	artilc := models.Article{}
	artilc.Id = id
	_, err := o.Delete(&artilc)
	if err != nil {
		beego.Info("Delete err:", err)
		this.Redirect("Article/ShowArticle", 302)
		return
	}

	this.Layout = "layout.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["contentHead"] = "文章详情页"
	this.Redirect("ShowArticle", 302)
}

func (this *ArticleController) ShowUpdate() {
	id, _ := this.GetInt("id")
	if id == 0 {
		beego.Info("链接故障")
		return
	}
	o := orm.NewOrm()
	artilc := models.Article{}
	artilc.Id = id
	err := o.Read(&artilc)
	if err != nil {
		beego.Info("Update err:", err)
		this.Redirect("Article/ShowArticle", 302)
		return
	}
	this.Data["artilc"] = artilc
	this.Layout="layout.html"
	this.TplName = "update.html"
}

func (this *ArticleController) UpdateArticle() {
	id := this.GetString("Id")
	title := this.GetString("articleName")
	content := this.GetString("content")
	//articleType:=this.GetString("select")

	f, h, err := this.GetFile("uploadname")
	if h.Filename == "" {
		beego.Info("err:", err)
		this.TplName = "Article/update.html"
		return
	}
	if err != nil {
		beego.Info("err:", err)
		this.TplName = "Article/update.html"
		return
	}
	defer f.Close()
	//1.文件格式
	ext := path.Ext(h.Filename)
	if ext != ".jpg" && ext != ".png" && ext != ".jpeg" {
		beego.Info("path.Ext 上传格式不正确", err)
		this.TplName = "Article/update.html"
		return
	}
	//2.文件大小
	if h.Size > 409600 {
		beego.Info("上传文件太大，不允许上传！ ", err)
		this.TplName = "Article/update.html"
		return
	}
	//3不能重名F:\GoCode\src\news\static\img
	//filename:=time.Now().Format("2006-01-02 15:04:05")
	filename := strconv.Itoa(time.Now().Nanosecond())
	err = this.SaveToFile("uploadname", "./static/img/"+filename+ext)
	if err != nil {
		beego.Info("SaveToFile err:", err)
		this.TplName = "Article/update.html"
		return
	}

	o := orm.NewOrm()
	artice := models.Article{}
	artice.Id, _ = strconv.Atoi(id)
	artice.Title = title
	artice.Content = content
	artice.Img = "./static/img/" + filename + ext

	_, err = o.Update(&artice)
	if err != nil {
		beego.Info("Update err:", err)
		this.TplName = "Article/add.html"
		return
	}
	this.Redirect("Article/ShowArticle", 302)

}
