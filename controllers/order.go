package controllers

import (
	"fanyouApi/service"
	"fmt"
	"github.com/bitly/go-simplejson"
)

type OrderController struct {
	BaseController
}

func (this *OrderController) CreateOrder() {
	username := this.Data["username"].(string)
	res := make(map[string]interface{})
	params, errJson := simplejson.NewJson(this.Ctx.Input.RequestBody)
	if errJson != nil {
		fmt.Println(errJson)
		res["error"] = errJson.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "json dump wrong", res}
		this.ServeJSON()
		this.StopRun()
	}
	foodId, errId := params.Get("food_id").Int()
	if errId != nil {
		fmt.Println(errId)
		res["error"] = errId.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "json dump food_id wrong", res}
		this.ServeJSON()
		this.StopRun()
	}
	orderId, errOr := service.CreateOrder(username, int(foodId))
	if errOr != nil {
		fmt.Println(errOr)
		res["error"] = errOr.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "create order failed . ", res}
		this.ServeJSON()
		this.StopRun()
	}
	res["order_id"] = orderId
	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = ReturnInfo{0, "create order success", res}
	this.ServeJSON()
	this.StopRun()
}

func (this *OrderController) GetOrderList() {
	username := this.Data["username"].(string)
	res := make(map[string]interface{})

	offset, errOff := this.GetInt("offset", 0)
	if errOff != nil {
		fmt.Println(errOff)
		res["error"] = errOff.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "get offset num failed . ", res}
		this.ServeJSON()
		this.StopRun()
	}
	limit, errLim := this.GetInt("limit", 0)
	if errLim != nil {
		fmt.Println(errLim)
		res["error"] = errLim.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "get limit num failed . ", res}
		this.ServeJSON()
		this.StopRun()
	}
	waitOrNot, errWait := this.GetInt("wait_or_not", 0)
	if errWait != nil {
		fmt.Println(errWait)
		res["error"] = errWait.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "get waitOrNot num failed . ", res}
		this.ServeJSON()
		this.StopRun()
	}
	total, orderList, err := service.GetOrderList(offset, limit, username, waitOrNot)
	if err != nil {
		res["error"] = err.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "get order list failed . ", res}
		this.ServeJSON()
		this.StopRun()
	}
	res["total"] = total
	res["order_list"] = orderList
	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = ReturnInfo{0, "get order list success", res}
	this.ServeJSON()
	this.StopRun()
}

func (this *OrderController) GetOrDenyOrder() {
	//username := this.Data["username"].(string)
	res := make(map[string]interface{})
	params, errJson := simplejson.NewJson(this.Ctx.Input.RequestBody)
	if errJson != nil {
		fmt.Println(errJson)
		res["error"] = errJson.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "json dump wrong", res}
		this.ServeJSON()
		this.StopRun()
	}
	orderId, errId := params.Get("order_id").Int()
	if errId != nil {
		fmt.Println(errId)
		res["error"] = errId.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "json dump order_id wrong", res}
		this.ServeJSON()
		this.StopRun()
	}
	changeStatus, errCha := params.Get("change_status").String()
	if errId != nil {
		fmt.Println(errCha)
		res["error"] = errCha.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "json dump change_status wrong", res}
		this.ServeJSON()
		this.StopRun()
	}
	err := service.ChangeOrderStatus(orderId, changeStatus)
	if err != nil {
		res["error"] = err.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "change order status failed . ", res}
		this.ServeJSON()
		this.StopRun()
	}
	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = ReturnInfo{0, "change order status success", res}
	this.ServeJSON()
	this.StopRun()
}

func (this *OrderController) DeleteHistory() {
	username := this.Data["username"].(string)
	err := service.DeleteHistoryByName(username)
	res := make(map[string]interface{})
	if err != nil {
		res["error"] = err.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "delete history failed . ", res}
		this.ServeJSON()
		this.StopRun()
	}
	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = ReturnInfo{0, "delete history success . ", res}
	this.ServeJSON()
	this.StopRun()
}
