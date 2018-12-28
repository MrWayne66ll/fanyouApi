package service

import (
	"errors"
	"fanyouApi/models"
	"time"
)

func CreateFood(username string, foodName string, foodType string, foodData string, comment string)(int,error){
	_,errUser:=models.GetUserByName(username)
	if errUser!=nil{
		return -1,errUser
	}
	if foodName==""{
		return -1,errors.New("food_name cannot be empty . ")
	}
	if foodType==""{
		return -1,errors.New("food_type cannot be empty . ")
	}
	if foodData==""{
		return -1,errors.New("food_date cannot be empty . ")
	} else {
		timeStr:= foodData+" 12:59:59 PM"
		tm,errTm:=time.Parse("2006-01-02 03:04:05 PM",timeStr)
		if errTm!=nil{
			return -1,errTm
		}
		timeNow := int(time.Now().Unix())
		if int(tm.Unix())<timeNow{
			return -1,errors.New("food_date cannot be early . ")
		}
		}
	foodId,errMod:=models.CreateFood(username,foodName,foodType,foodData,comment)
	if errMod!=nil{
		return -1,errMod
	}
	return foodId,nil
	}
