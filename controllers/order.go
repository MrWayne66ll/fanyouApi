package controllers

import (
	"fanyouApi/service"
	"fmt"
	"github.com/bitly/go-simplejson"
)

type OrderController struct {
	BaseController
}

func (this *OrderController) CreateOrder(){
	username := this.Data["username"].(string)
	res := make(map[string]interface{})
	params,errJson := simplejson.NewJson(this.Ctx.Input.RequestBody)
	if errJson!=nil{
		fmt.Println(errJson)
		res["error"] = errJson.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "json dump wrong", res}
		this.ServeJSON()
		this.StopRun()
	}
	foodId,errId:=params.Get("food_id").Int()
	if errId!=nil{
		fmt.Println(errId)
		res["error"] = errId.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "json dump food_id wrong", res}
		this.ServeJSON()
		this.StopRun()
	}
	orderId,errOr:=service.CreateOrder(username,int(foodId))
	if errOr!=nil{
		fmt.Println(errOr)
		res["error"] = errOr.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "create order failed . ", res}
		this.ServeJSON()
		this.StopRun()
	}
	res["order_id"] = orderId
	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = ReturnInfo{1000, "create order success", res}
	this.ServeJSON()
	this.StopRun()
}