package models

import "github.com/astaxie/beego/orm"

const(
	ORDER_WAITING 	= "wait"		// 等待状态
	ORDER_GET 		= "get"			// 单子已被获取
	ORDER_INVALID 	= "invalid"		// 单子失效
	ORDER_DENY 		= "deny"		// 反悔单子
)

type Order struct {
	Id 			int
	FoodId		int
	CatchUserId	int
	CatchTime	string
	GetTime		string
	Status		string
}


func init(){
	orm.RegisterModel(new(Order))
}