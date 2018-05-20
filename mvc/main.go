package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
)

var db gorm.DB

func main() {
	connect()
	rooter(goji.DefaultMux)
	goji.Serve()
}

func rooter(m *web.Mux) http.Handler {
	m.Use(SuperSecure)

	user := web.New()
	goji.Handle("/user/*", user)
	user.Use(middleware.SubRouter)
	user.Get("/index", UserIndex)
	user.Get("/new", UserNew)
	user.Post("/new", UserCreate)
	user.Get("/edit/:id", UserEdit)
	user.Post("/update/:id", UserUpdate)
	user.Get("/delete/:id", UserDelete)

	return m
}

func connect() {
	yml, err := ioutil.ReadFile("conf/db.yml")
	if err != nil {
		panic(err)
	}

	t := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(yml), &t)
	if err != nil {
		panic(err)
	}
	env := os.Getenv("GOJIENV")
	if env == "" {
		env = "development"
	}

	conn := t[env].(map[interface{}]interface{})
	db, err = gorm.Open("mysql", conn["user"].(string)+":"+conn["password"].(string)+"@/"+conn["db"].(string)+"?charset=utf8&parseTime=True")
	if err != nil {
		panic(err)
	}
}
