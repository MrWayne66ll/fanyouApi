package service

import (
	"errors"
	"fanyouApi/models"
	"github.com/astaxie/beego/orm"
)

// 创建一个order单
func CreateOrder(username string, foodId int)(int,error){
	_,errFood:=models.GetFoodById(foodId)
	if errFood!=nil{
		return -1,errFood
	}
	orderId,errOr:=models.CreateOrder(username,foodId)
	if errOr!=nil{
		return -1,errOr
	}
	return orderId,nil
}

// 获取个人的单子列表，1表示未结束的order单，2表示已经结束的单子
func GetOrderList(username string,waitOrNot int)(int,[]orm.Params,error){
	if waitOrNot!=1&&waitOrNot!=2{
		return -1,[]orm.Params{},errors.New("waitOrNot input wrong , must be 1 or 2 . ")
	}
	total,orderList,errList:=models.GetOrderList(username,waitOrNot)
	if errList!=nil{
		return -1,[]orm.Params{},errList
	}
	return total,orderList,nil
	}