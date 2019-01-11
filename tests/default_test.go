package test

import (
	_ "fanyouApi/routers"
	"fanyouApi/service"
	"github.com/astaxie/beego/orm"
	"testing"

	//"net/http"
	//"net/http/httptest"
	"path/filepath"
	"runtime"
	//"testing"

	"github.com/astaxie/beego"
	//. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestBeego is a sample to run an endpoint test
//func TestBeego(t *testing.T) {
//	r, _ := http.NewRequest("GET", "/", nil)
//	w := httptest.NewRecorder()
//	beego.BeeApp.Handlers.ServeHTTP(w, r)
//
//	beego.Trace("testing", "TestBeego", "Code[%d]\n%s", w.Code, w.Body.String())
//
//	Convey("Subject: Test Station Endpoint\n", t, func() {
//		Convey("Status Code Should Be 200", func() {
//			So(w.Code, ShouldEqual, 200)
//		})
//		Convey("The Result Should Not Be Empty", func() {
//			So(w.Body.Len(), ShouldBeGreaterThan, 0)
//		})
//	})
//}

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

func TestDeleteHistoryByName(t *testing.T) {
	initDb()
	username := "fengchuanling"
	err := service.DeleteHistoryByName(username)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("success")
	}
}
