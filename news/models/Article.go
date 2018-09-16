package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id int `orm:"pk;auto"`
	Title string `orm:"size(20)"`
	Time time.Time `orm:"auto_now;type(datetime)"`
	Count int `orm:"default(0);null"`
	Content string `orm:"size(500)"`
	Img string  `orm:"size(100)"`
	ArticleType *ArticleType `orm:"rel(fk)"`
	Users []*User `orm:"reverse(many)"`
}

func init()  {
	orm.RegisterModel(new(Article))
}
