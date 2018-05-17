package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/picobank/instruments/routers"

	_ "github.com/lib/pq"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {

	orm.RegisterDriver("postgres", orm.DRPostgres)
	parts := []interface{}{
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	}
	connectionString := fmt.Sprintf("user='%s' password='%s' host='%s' dbname='%s' sslmode=disable", parts...)
	log.Println("Connection string: ", connectionString)
	orm.RegisterDataBase("default", "postgres", connectionString)
	orm.SetMaxIdleConns("default", 10)
	orm.SetMaxOpenConns("default", 10)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
