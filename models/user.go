package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id			int
	Username	string
	UsernameCn	string
	Email 		string
	Floor		string
	Shelf		string
	Active		int
	CreateTime	string
}

func init(){
	orm.RegisterModel(new(User))
}

// 根据username获取userid
func GetUserByName(username string)(int,error){
	o := orm.NewOrm().QueryTable("user")
	user := User{}
	err := o.Filter("username",username).One(&user)
	if err!=nil{
		return -1,err
	}
	return user.Id,nil
}

// 创建新用户
func CreateUser(username string,usernameCn string,floor string,shelf string)(int,error){
	user := User{}
	user.Username = username
	user.UsernameCn = usernameCn
	user.Floor = floor
	user.Shelf = shelf
	user.Email = username + "@pinduoduo.com"
	user.Active = 1
	user.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	o := orm.NewOrm()
	id,err := o.Insert(&user)
	if err!=nil{
		return -1,err
	}
	return int(id),nil
}