package controllers

import (
	"fanyouApi/models"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

type ReturnInfo struct {
	Code    int         `json:"error_code"`
	Message string      `json:"error_msg"`
	Data    interface{} `json:"data"`
}

func (this *BaseController) Prepare() {
	headerInfo := this.Ctx.Request.Header
	if headerInfo.Get("Username") == "" {
		this.Ctx.Output.SetStatus(403)
		this.Data["json"] = ReturnInfo{1000, "head user cannot be empty . ", map[string]string{}}
		this.ServeJSON()
		return
	}
	this.Data["username"] = headerInfo["Username"][0]
	_, err := models.GetUserByName(this.Data["username"].(string))
	if err != nil {
		this.Ctx.Output.SetStatus(403)
		this.Data["json"] = ReturnInfo{1000, "operator not found", map[string]string{}}
		this.ServeJSON()
		return
	}
}
