package service

import "fanyouApi/models"

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
