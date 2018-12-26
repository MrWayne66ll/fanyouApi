package test

import (
	"fanyouApi/models"
	"github.com/astaxie/beego/orm"
	"testing"
)

func initDb(){
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
func TestCreateUser(t *testing.T){
	initDb()
	username := "fengchuanling"
	username_cn := "枫川棂"
	floor := "28"
	shelf := "f3"
	id,err := models.CreateUser(username,username_cn,floor,shelf)
	if err!=nil{
		t.Error(err)
	} else {
		t.Log(id)
	}
	}

// 测试根据name获取用户
func TestGetUserByName(t *testing.T){
	initDb()
	username := "fengchuanling"
	id,err:=models.GetUserByName(username)
	if err!=nil{
		t.Error(err)
	} else {
		t.Log(id)
	}
	}