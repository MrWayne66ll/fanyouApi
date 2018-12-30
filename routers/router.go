package routers

import (
	"fanyouApi/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api/v0",
		beego.NSNamespace("/food",
			beego.NSRouter("/foodlist", &controllers.FoodController{},"get:GetFoodList"),
			beego.NSRouter("/create", &controllers.FoodController{},"post:CreateFood"),
			),
		beego.NSNamespace("/order",
			beego.NSRouter("/create", &controllers.OrderController{},"post:CreateOrder"),
			beego.NSRouter("/orderlist", &controllers.OrderController{},"get:GetOrderList"),
			),
	)
	beego.AddNamespace(ns)
}
