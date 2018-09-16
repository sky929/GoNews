package main

import (
	"github.com/astaxie/beego"
	_ "news/routers"
)

func main() {
	beego.AddFuncMap("PrePage",HandlePrePage)
	beego.AddFuncMap("NextPage",HandleNextPage)
	beego.Run()
}

func HandlePrePage(datd int) int {
	page := datd - 1
	return page
}
func HandleNextPage(datd int) int {
	page := datd + 1
	return page
}
