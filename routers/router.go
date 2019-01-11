package routers

import (
	"fanyouApi/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	ns := beego.NewNamespace("/api/v0",
		beego.NSNamespace("/food",
			beego.NSRouter("/foodlist", &controllers.FoodController{}, "get:GetFoodList"),
			beego.NSRouter("/create", &controllers.FoodController{}, "post:CreateFood"),
		),
		beego.NSNamespace("/order",
			beego.NSRouter("/create", &controllers.OrderController{}, "post:CreateOrder"),
			beego.NSRouter("/orderlist", &controllers.OrderController{}, "get:GetOrderList"),
			beego.NSRouter("/getdeny", &controllers.OrderController{}, "post:GetOrDenyOrder"),
			beego.NSRouter("/deletehistory", &controllers.OrderController{}, "get:DeleteHistory"),
		),
	)
	beego.AddNamespace(ns)
}
