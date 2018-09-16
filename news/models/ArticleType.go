package models

import "github.com/astaxie/beego/orm"

type ArticleType struct {
	Id int   `orm:"pk;aout`
	TypeName string `orm:"size(20)"`
	Articles []*Article `orm:"reverse(many);null"`
}
func init()  {
	orm.RegisterModel(new(ArticleType))
}
