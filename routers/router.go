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
	)
	beego.AddNamespace(ns)
}
