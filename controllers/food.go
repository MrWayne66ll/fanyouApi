package controllers

import (
	"fanyouApi/service"
	"github.com/bitly/go-simplejson"
	"strings"
)

type FoodController struct {
	BaseController
}

func (this *FoodController) GetFoodList() {
	offset,errOffset :=this.GetInt("offset",0)
	if errOffset!=nil{
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "get input offset wrong", map[string]string{}}
		this.ServeJSON()
		this.StopRun()
	}
	limit,errLimit := this.GetInt("limit",0)
	if errLimit!=nil{
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "get input offset wrong", map[string]string{}}
		this.ServeJSON()
		this.StopRun()
	}
	username := this.GetString("search_username","")
	foodType:= this.GetString("food_type","")
	startTime := this.GetString("start_time","")
	endTime:= this.GetString("end_time","")
	total,foodList,err:=service.GetFoodList(offset,limit,username,foodType,startTime,endTime)
	if err!=nil{
		res := make(map[string]interface{})
		res["error"] = err.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "get food list wrong", res}
		this.ServeJSON()
		this.StopRun()
	}
	res := make(map[string]interface{})
	res["total"] = total
	res["food_list"] = foodList
	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = ReturnInfo{0, "get food list success", res}
	this.ServeJSON()
	this.StopRun()
	}

func (this *FoodController) CreateFood(){
	res := make(map[string]interface{})
	params,errJson := simplejson.NewJson(this.Ctx.Input.RequestBody)
	if errJson!=nil{
		res["error"] = errJson.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "json dump wrong", res}
		this.ServeJSON()
		this.StopRun()
	}
	foodName,errFoodName:=params.Get("food_name").String()
	if errFoodName!=nil{
		res["error"] = errFoodName.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "json dump food_name wrong", res}
		this.ServeJSON()
		this.StopRun()
	}
	foodName = strings.Replace(foodName," ","",-1)
	foodType,errFoodType:= params.Get("food_type").String()
	if errFoodType!=nil{
		res["error"] = errFoodType.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "json dump food_type wrong", res}
		this.ServeJSON()
		this.StopRun()
	}
	foodType = strings.Replace(foodType," ","",-1)
	food_date,errFoodDate:=params.Get("food_date").String()
	if errFoodDate!=nil{
		res["error"] = errFoodDate.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "json dump food_date wrong", res}
		this.ServeJSON()
		this.StopRun()
	}
	food_date = strings.Replace(food_date," ","",-1)
	comment,errCom:=params.Get("comment").String()
	if errCom!=nil{
		res["error"] = errCom.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "json dump comment wrong", res}
		this.ServeJSON()
		this.StopRun()
	}
	username := this.Data["username"].(string)
	foodId,errSer:=service.CreateFood(username,foodName,foodType,food_date,comment)
	if errSer!=nil{
		res["error"] = errSer.Error()
		this.Ctx.Output.SetStatus(500)
		this.Data["json"] = ReturnInfo{1000, "create food failed . ", res}
		this.ServeJSON()
		this.StopRun()
	}
	res["food_id"] = foodId
	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = ReturnInfo{0, "create food success . ", res}
	this.ServeJSON()
	this.StopRun()
	}