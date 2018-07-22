package controller

import (
	"net/http"
	"io"
	"github.com/astaxie/beego/logs"
	"beegoweb/model"
	"encoding/json"
	"strconv"
)

func Test(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))

	user := model.User{Name: name, Age: age}
	userJson , err := json.Marshal(user)
	if err != nil {
		logs.Info("json格式转换错误")
		return
	}
	io.WriteString(w, string(userJson))
}

