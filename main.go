package main

import (
	"github.com/astaxie/beego"
	"beegoweb/controller"
)

func main()  {
	beego.Router("/", &beegoweb.WebController{})
	beego.SetStaticPath("/views", "views")
	beego.Run()
}

