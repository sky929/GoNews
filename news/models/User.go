package models

import (
	_ "github.com/Go-SQL-Driver/mysql"
	"github.com/astaxie/beego/orm"
)
type User struct {
	Id int
	UserName string
	Passwd string
	Articles []*Article`orm:"rel(m2m)"`

}
func init(){
	orm.RegisterDataBase("default","mysql","root:123456@tcp(192.168.88.134:3306)/newweb")
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default",false,true)
}

