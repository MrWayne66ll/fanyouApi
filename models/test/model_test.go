package test

import (
	"fanyouApi/models"
	"github.com/astaxie/beego/orm"
	"testing"
)

func initDb() {
	dbHost := "192.168.17.128"
	dbPort := "3306"
	dbUser := "root"
	dbPassword := "123123"
	db := "fanyou"

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + db + "?charset=utf8"

	orm.RegisterDataBase("default", "mysql", dsn)
	orm.SetMaxIdleConns("default", 120)
	orm.SetMaxOpenConns("default", 120)
}

// 测试创建用户
func TestCreateUser(t *testing.T) {
	initDb()
	username := "fengchuanling"
	username_cn := "枫川棂"
	floor := "28"
	shelf := "f3"
	id, err := models.CreateUser(username, username_cn, floor, shelf)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(id)
	}
}

// 测试根据name获取用户
func TestGetUserByName(t *testing.T) {
	initDb()
	username := "fengchuanling"
	id, err := models.GetUserByName(username)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(id)
	}
}

// 测试创建食物
func TestCreateFood(t *testing.T) {
	initDb()
	username := "fengchuanling"
	foodName := "烤肉饭"
	foodType := "breakfast"
	foodDate := "2018-12-27"
	i, err := models.CreateFood(username, foodName, foodType, foodDate, "")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(i)
	}
}

// 测试获取所有食物列表
func TestGetFoodList(t *testing.T) {
	initDb()
	offset := 0
	limit := 0
	username := ""
	foodType := ""
	startTime := ""
	endTime := ""
	total, foodList, err := models.GetFoodList(offset, limit, username, foodType, startTime, endTime)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(total)
		for i, v := range foodList {
			t.Log(i)
			t.Log(v)
		}
	}
}

// 测试下架食物
func TestInActiveFood(t *testing.T) {
	initDb()
	foodId := 1
	err := models.InActiveFood(foodId)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("success")
	}
}
