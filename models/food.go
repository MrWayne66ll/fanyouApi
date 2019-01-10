package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

const (
	FOOD_STATUS_RELEASE = "release" // 已经发布，还未被抢
	FOOD_STATUS_Catch   = "catch"   // 食物被抢
	FOOD_STATUS_GET     = "get"     // 食物已经领取
)
const (
	FOODTYPE_BRE = "breakfast"
	FOODTYPE_LUN = "lunch"
	FOODTYPE_DIN = "dinner"
	FOODTYPE_NIG = "nightingale"
)

type Food struct {
	Id          int
	UserId      int
	FoodName    string
	FoodDate    string
	Status      string
	ReleaseTime string
	GetTime     string
	FoodType    string
	Comment     string
	Active      int
}

func init() {
	orm.RegisterModel(new(Food))
}

// 创建一个饭
func CreateFood(username string, foodName string, foodType string, foodData string, comment string) (int, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	var food Food
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
		return -1, errors.New("Get wrong food type !")
	}
	user, errUser := GetUserByName(username)
	if errUser != nil {
		return -1, errUser
	}
	food.UserId = user.Id
	food.Status = FOOD_STATUS_RELEASE
	food.ReleaseTime = timestamp
	food.GetTime = timestamp
	food.FoodName = foodName
	food.FoodDate = foodData
	if comment != "" {
		food.Comment = comment
	} else {
		food.Comment = "快来取我吧~"
	}
	food.Active = 1
	o := orm.NewOrm()
	foodId, errFood := o.Insert(&food)
	if errFood != nil {
		return -1, errFood
	}
	return int(foodId), nil
}

// 获取食物列表
func GetFoodList(offset int, limit int, username string, foodType string, startTime string, endTime string) (int, []orm.Params, error) {
	o := orm.NewOrm()
	sql := `select * from food where active=1 and status="release"`
	if username != "" {
		user,errUser:=GetUserByName(username)
		if errUser!=nil{
			return -1,[]orm.Params{},errUser
		}
		sql = sql + " and user_id=" + strconv.Itoa(user.Id)
	}
	switch foodType {
	case FOODTYPE_BRE:
		sql = sql + " and food_type=" + `"`+FOODTYPE_BRE+`"`
	case FOODTYPE_LUN:
		sql = sql + " and food_type=" + `"`+FOODTYPE_LUN+`"`
	case FOODTYPE_DIN:
		sql = sql + " and food_type=" + `"`+FOODTYPE_DIN+`"`
	case FOODTYPE_NIG:
		sql = sql + " and food_type=" + `"`+FOODTYPE_NIG+`"`
	}
	if startTime != "" && endTime != "" {
		sql = sql + " and (food_date between " + `"` + startTime + `"` + " and " + `"` + endTime + `"` + ")"
	} else {
		if startTime != "" {
			sql = sql + " and food_date>" + `"` + startTime + `"`
		}
		if endTime != "" {
			sql = sql + " and food_date<" + `"` + endTime + `"`
		}
	}
	if limit > 0 {
		sql = sql + " limit " + strconv.Itoa(offset) + "," + strconv.Itoa(limit)
	}
	sql = `select 
			user.username as username,
			user.username_cn as username_cn,
			user.floor as user_floor,
			f.food_name,
			f.food_type,
			case f.food_type
				when "breakfast" then "早餐"
				when "lunch" then "午餐"
				when "dinner" then "晚餐"
				when "nightingale" then "夜宵"
				else "other"
			end
			as food_type_cn,
			f.food_date,
			f.comment,
			f.status as food_status,
			f.id as id
			from (` + sql + `) as f left join user on f.user_id=user.id where user.active = 1 order by id desc`
	var foodList []orm.Params
	total, err := o.Raw(sql).Values(&foodList)
	if err != nil {
		return 0, foodList, err
	}
	return int(total), foodList, nil
}

// 获取一个食物详情
func GetFoodById(foodId int) (Food, error) {
	food := Food{}
	food.Id = foodId
	o := orm.NewOrm()
	err := o.Read(&food)
	if err != nil {
		return Food{}, err
	}
	return food, nil
}

// 用户下架一个饭
func InActiveFood(foodId int) error {
	o := orm.NewOrm()
	food := Food{}
	food.Id = foodId
	err := o.Read(&food)
	if err != nil {
		return err
	}
	food.Active = 0
	_, errUp := o.Update(&food, "active")
	if errUp != nil {
		return errUp
	}
	return nil
}

// 修改食物的状态
func ChangeFoodStatus(foodId int, changeStatus string,timeStamp string) error {
	food := Food{}
	food.Id = foodId
	o := orm.NewOrm()
	errRe := o.Read(&food)
	if errRe != nil {
		return errRe
	}
	switch changeStatus {
	case FOOD_STATUS_RELEASE:
		food.Status = FOOD_STATUS_RELEASE
	case FOOD_STATUS_Catch:
		food.Status = FOOD_STATUS_Catch
	case FOOD_STATUS_GET:
		food.Status = FOOD_STATUS_GET
	default:
		return errors.New("wrong food status input . ")
	}
	if timeStamp != ""{
		food.GetTime = timeStamp
		_, errUp := o.Update(&food,"status","get_time")
		if errUp != nil {
			return errUp
		}
	}
	_, errUp := o.Update(&food,"status")
	if errUp != nil {
		return errUp
	}
	return nil
}
