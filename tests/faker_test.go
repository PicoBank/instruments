package test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/picobank/instruments/models"
	_ "github.com/picobank/instruments/routers"

	_ "github.com/lib/pq"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"github.com/bxcodec/faker"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))

	if beego.BConfig.RunMode == beego.DEV {
		orm.Debug = true
	}

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

	beego.TestBeegoInit(apppath)
}

// This test generates a stack overflow (probably due to self references)
func TestFaker(t *testing.T) {
	a := models.Instrument{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)
}
