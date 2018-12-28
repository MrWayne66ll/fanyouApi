package main

import (
	"fanyouApi/models"
	_ "fanyouApi/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	beego.Run()
}
