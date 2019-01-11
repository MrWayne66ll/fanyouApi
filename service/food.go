package service

import (
	"errors"
	"fanyouApi/models"
	"github.com/astaxie/beego/orm"
	"time"
)

func GetFoodList(offset int, limit int, username string, foodType string, startTime string, endTime string) (int, []orm.Params, error) {
	if startTime != "" && endTime != "" {
		startStr := startTime + " 00:00:01 AM"
		tmStart, errTmStart := time.Parse("2006-01-02 03:04:05 PM", startStr)
		if errTmStart != nil {
			return -1, []orm.Params{}, errTmStart
		}
		endStr := endTime + " 12:59:59 PM"
		tmEnd, errTmEnd := time.Parse("2006-01-02 03:04:05 PM", endStr)
		if errTmEnd != nil {
			return -1, []orm.Params{}, errTmEnd
		}
		if int(tmStart.Unix()) >= int(tmEnd.Unix()) {
			return -1, []orm.Params{}, errors.New("start_time > end_time , it's wrong ! ")
		}

	}
	total, foodList, err := models.GetFoodList(offset, limit, username, foodType, startTime, endTime)
	if err != nil {
		return -1, []orm.Params{}, err
	}
	return total, foodList, nil
}

// 创建food
func CreateFood(username string, foodName string, foodType string, foodData string, comment string) (int, error) {
	_, errUser := models.GetUserByName(username)
	if errUser != nil {
		return -1, errUser
	}
	if foodName == "" {
		return -1, errors.New("food_name cannot be empty . ")
	}
	if foodType == "" {
		return -1, errors.New("food_type cannot be empty . ")
	}
	if foodData == "" {
		return -1, errors.New("food_date cannot be empty . ")
	} else {
		timeStr := foodData + " 11:59:59 PM"
		tm, errTm := time.Parse("2006-1-2 03:04:05 PM", timeStr)
		if errTm != nil {
			return -1, errTm
		}
		timeNow := int(time.Now().Unix())
		if int(tm.Unix()) < timeNow {
			return -1, errors.New("food_date cannot be early . ")
		}
	}
	foodId, errMod := models.CreateFood(username, foodName, foodType, foodData, comment)
	if errMod != nil {
		return -1, errMod
	}
	return foodId, nil
}
