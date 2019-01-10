package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

const (
	ORDER_WAITING = "wait"    // 等待状态
	ORDER_GET     = "get"     // 单子已被获取
	ORDER_INVALID = "invalid" // 单子失效
	ORDER_DENY    = "deny"    // 反悔单子
)

type OrderFood struct {
	Id          int
	FoodId      int
	CatchUserId int
	CatchTime   string
	GetTime     string
	Status      string
	Active      int
}

func init() {
	orm.RegisterModel(new(OrderFood))
}

// 创建一个抢单
func CreateOrder(username string, foodId int) (int, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	user, errUser := GetUserByName(username)
	if errUser != nil {
		return -1, errUser
	}
	food, errFood := GetFoodById(foodId)
	if errFood != nil {
		return -1, errFood
	}
	if food.Active != 1 {
		return -1, errors.New("this food is already deleted . ")
	}
	if food.Status != FOOD_STATUS_RELEASE {
		return -1, errors.New("this food has been ordered already . ")
	}
	order := OrderFood{}
	order.CatchUserId = user.Id
	order.FoodId = foodId
	order.CatchTime = timestamp
	order.GetTime = timestamp
	order.Status = ORDER_WAITING
	order.Active = 1
	o := orm.NewOrm()
	orderId, errId := o.Insert(&order)
	if errId != nil {
		return -1, errId
	}
	errChang := ChangeFoodStatus(foodId, FOOD_STATUS_Catch,"")
	if errChang != nil {
		InActiveOrder(int(orderId))
		return -1, errChang
	}
	return int(orderId), nil
}

// 获取单个抢单
func GetOrderById(orderId int) (OrderFood, error) {
	order := OrderFood{}
	order.Id = orderId
	o := orm.NewOrm()
	err := o.Read(&order)
	if err != nil {
		return OrderFood{}, err
	}
	return order, nil
}

// 修改抢单状态
func ChangeOrderStatus(orderId int, changeStatus string, timeStamp string) error {
	order, errOr := GetOrderById(orderId)
	if errOr != nil {
		return errOr
	}
	if order.Status == changeStatus {
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
	if timeStamp != "" {
		order.GetTime = timeStamp
		o := orm.NewOrm()
		_, err := o.Update(&order, "status","get_time")
		if err != nil {
			return err
		}
	}
	o := orm.NewOrm()
	_, err := o.Update(&order, "status")
	if err != nil {
		return err
	}
	return nil
}

// 关闭一个抢单
func InActiveOrder(orderId int) error {
	o := orm.NewOrm()
	order := OrderFood{}
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

// 获取orderlist
func GetOrderList(username string,waitOrNot int)(int,[]orm.Params,error){
	user,errUser:=GetUserByName(username)
	if errUser!=nil{
		return -1,[]orm.Params{},errUser
	}
	o:= orm.NewOrm()
	sql := `select * from order_food where active=1`
	if username !=""{
		sql = sql + ` and catch_user_id="` + strconv.Itoa(user.Id) + `"`
	}
	switch waitOrNot {
	case 1:
		sql = sql + ` and status="wait"`
	case 2:
		sql = sql + ` and status!="wait"`
	}
	sql = `select 
		o.id as order_id,
		o.status as order_status,
		case o.status
			when "wait" then "等待获取"
			when "get" then "已获取"
			when "invalid" then "失效"
			when "deny" then "丢弃"
			else "other"
		end
		as order_status_cn,
		o.catch_time,
		food.food_name,
		food.food_date,
		food.food_type,
		case food.food_type
			when "breakfast" then "早餐"
			when "lunch" then "午餐"
			when "dinner" then "晚餐"
			when "nightingale" then "夜宵"
			else "other"
		end
		as food_type_cn,
		user.username_cn,
		user.floor,user.shelf 
		from (` + sql + `) as o 
		left join food on food_id=food.id
		left join user on food.user_id=user.id 
		where user.active=1 and food.active=1`
	var orderList []orm.Params
	total,errSql:=o.Raw(sql).Values(&orderList)
	if errSql!=nil{
		return -1,orderList,errSql
	}
	return int(total),orderList,nil
}
