package beegoweb

import "github.com/astaxie/beego"

type WebController struct {
	beego.Controller
}

func (webController *WebController) Get()  {
	webController.Ctx.WriteString("go web")
}