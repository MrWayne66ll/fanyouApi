package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func Init() {
	//dbHost := beego.AppConfig.String("db.host")
	//dbPort := beego.AppConfig.String("db.port")
	//dbUser := beego.AppConfig.String("db.user")
	//dbPassword := beego.AppConfig.String("db.password")
	//db := beego.AppConfig.String("db.name")
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
