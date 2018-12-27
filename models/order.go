package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	ORDER_WAITING = "wait"    // 等待状态
	ORDER_GET     = "get"     // 单子已被获取
	ORDER_INVALID = "invalid" // 单子失效
	ORDER_DENY    = "deny"    // 反悔单子
)

type Order struct {
	Id          int
	FoodId      int
	CatchUserId int
	CatchTime   string
	GetTime     string
	Status      string
	Active		int
}

func init() {
	orm.RegisterModel(new(Order))
}

// 创建一个抢单
func CreateOrder(username string,foodId int)(int,error){
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	user,errUser:=GetUserByName(username)
	if errUser !=nil{
		return -1,errUser
	}
	food,errFood := GetFoodById(foodId)
	if errFood!=nil{
		return -1,errFood
	}
	if food.Active != 1{
		return -1,errors.New("this food is already deleted . ")
	}
	if food.Status != FOOD_STATUS_RELEASE{
		return -1,errors.New("this food has been ordered already . ")
	}
	order := Order{}
	order.CatchUserId = user.Id
	order.FoodId = foodId
	order.CatchTime = timestamp
	order.GetTime = timestamp
	order.Status = ORDER_WAITING
	o := orm.NewOrm()
	orderId,errId:=o.Insert(&order)
	if errId !=nil{
		return -1,errId
	}
	errChang:=ChangeFoodStatus(foodId,FOOD_STATUS_Catch)
	if errChang!=nil{
		InActiveOrder(int(orderId))
		return -1,errChang
	}
	return int(orderId),nil
}

// 获取单个抢单
func GetOrderById(orderId int)(Order,error){
	order := Order{}
	order.Id = orderId
	o:= orm.NewOrm()
	err:=o.Read(&order)
	if err!=nil{
		return Order{},err
	}
	return order,nil
}

// 修改抢单状态
func ChangeOrderStatus(orderId int,changeStatus string)(error){
	order,errOr := GetOrderById(orderId)
	if errOr!=nil{
		return errOr
	}
	if order.Status == changeStatus{
		return errors.New("change status is the same status , get nothing to do .")
	}
	switch changeStatus {
	case ORDER_WAITING:
		order.Status = ORDER_WAITING
	case ORDER_DENY:
		order.Status = ORDER_DENY
	case ORDER_GET:
		order.Status = ORDER_GET
	case ORDER_INVALID:
		order.Status = ORDER_INVALID
	default:
		return errors.New("wrong order status input .")
	}
	o:= orm.NewOrm()
	_,err :=o.Update(&order,"status")
	if err!=nil{
		return err
	}
	return nil
}

// 关闭一个抢单
func InActiveOrder(orderId int) error {
	o := orm.NewOrm()
	order := Order{}
	order.Id = orderId
	err := o.Read(&order)
	if err != nil {
		return err
	}
	order.Active = 0
	_, errUp := o.Update(&order, "active")
	if errUp != nil {
		return errUp
	}
	return nil
}