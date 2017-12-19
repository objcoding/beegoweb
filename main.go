package main

import (
	"github.com/astaxie/beego"
	"beegoweb/controller"
)

func main()  {
	beego.SetStaticPath("/view", "static")
	beego.Router("/", &beegoweb.WebController{})

	beego.Run()
}

