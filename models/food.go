package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"time"
)

const(
	FOOD_STATUS_RELEASE = "release"		// 已经发布，还未被抢
	FOOD_STATUS_Catch = "catch"			// 食物被抢
	FOOD_STATUS_GET = "get"				// 食物已经领取
)
const(
	FOODTYPE_BRE = "breakfast"
	FOODTYPE_LUN = "lunch"
	FOODTYPE_DIN = "dinner"
	FOODTYPE_NIG = "nightingale"
)

type Food struct {
	Id 			int
	UserId		int
	FoodName	string
	Status 		string
	ReleaseTime	string
	GetTime		string
	FoodType	string
	Comment		string
}

func init(){
	orm.RegisterModel(new(Food))
}

// 创建一个饭
func CreateFood(username string,foodName string,foodType string,comment string)(int,error){
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	food := new(Food)
	switch foodType {
	case FOODTYPE_BRE:
		food.FoodType = FOODTYPE_BRE
	case FOODTYPE_LUN:
		food.FoodType = FOODTYPE_LUN
	case FOODTYPE_DIN:
		food.FoodType = FOODTYPE_DIN
	case FOODTYPE_NIG:
		food.FoodType = FOODTYPE_NIG
	default:
		return -1,errors.New("Get wrong food type !")
		}
	userId,errUser := GetUserByName(username)
	if errUser != nil{
		return -1,errUser
	}
	food.UserId = userId
	food.Status = FOOD_STATUS_RELEASE
	food.ReleaseTime = timestamp
	food.GetTime = timestamp
	food.FoodName = foodName
	food.Comment = comment
	o := orm.NewOrm()
	foodId,errFood := o.Insert(&food)
	if errFood!=nil{
		return -1,errFood
	}
	return int(foodId),nil
}