package service

import (
	"errors"
	"fanyouApi/models"
	"github.com/astaxie/beego/orm"
	"time"
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

// 修改单子状态（确认获取、反悔单子）
func ChangeOrderStatus(orderId int,changeStatus string)(error){
	order,errOrd := models.GetOrderById(orderId)
	if errOrd!=nil{
		return errOrd
	}
	// 如果为get，则修改order状态为get，并修改food为get，并且两者的get_time都为timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	if changeStatus == "get" {
		errCha := models.ChangeOrderStatus(orderId,changeStatus,timestamp)
		if errCha != nil{
			return errCha
		}
		errChaFood := models.ChangeFoodStatus(order.FoodId,models.FOOD_STATUS_GET,timestamp)
		if errCha!=nil{
			models.ChangeOrderStatus(orderId,models.ORDER_WAITING,"")
			return errChaFood
		}
	}
	// 如果为deny，则为反悔单子，修改order状态为deny，并且修改food为release，使得可以food可以继续抢夺
	if changeStatus == "deny" {
		errCha := models.ChangeOrderStatus(orderId,changeStatus,"")
		if errCha != nil{
			return errCha
		}
		errChaFood := models.ChangeFoodStatus(order.FoodId,models.FOOD_STATUS_RELEASE,"")
		if errCha!=nil{
			models.ChangeOrderStatus(orderId,models.ORDER_WAITING,"")
			return errChaFood
		}
	}
	return nil
}