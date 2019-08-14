package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"runtime"
	"bytes"
	"path/filepath"
	_ "pet/routers"
	"mime/multipart"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("sqlconn"))
	o := orm.NewOrm()
	o.Raw("TRUNCATE TABLE pet;").Exec() // DELETE ALL THE DATA FIRST
}

func TestGetAll(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/pet", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGetAll", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

func TestPost(t *testing.T) {
	var jsonStr = []byte(`{
		"name": "Hendrik",
		"age": 15
	}`)

	r, _ := http.NewRequest("POST", "/v1/pet", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestPost", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 201", func() {
			So(w.Code, ShouldEqual, 201)
		})
	})
}

func TestPut(t *testing.T) {
	var jsonStr = []byte(`{
		"name": "Kustanto",
		"age": 15
	}`)

	r, _ := http.NewRequest("PUT", "/v1/pet/1", bytes.NewBuffer(jsonStr))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestPut", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

func TestUploadImage(t *testing.T) {
	var buf bytes.Buffer
	mpw := multipart.NewWriter(&buf)
	mpw.CreateFormFile("image", "beego.png")
	mpw.Close()

	r, _ := http.NewRequest("POST", "/v1/pet/1/uploadImage", &buf)
	r.Header.Set("Content-Type", mpw.FormDataContentType())
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestUploadImage", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

func TestGetOne(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/pet/1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGetOne", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

func TestDelete(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/v1/pet/1", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestDelete", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}