package main

//go:generate sqlboiler -c database/sqlboiler.toml postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/picobank/instruments/routers"

	_ "github.com/lib/pq"
)

func init() {
	parts := []interface{}{
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	}
	connectionString := fmt.Sprintf("user='%s' password='%s' host='%s' dbname='%s' sslmode=disable", parts...)
	log.Println("Connection string: ", connectionString)
}

func main() {
	log.Println("Starting...")

	routers.Start()
}
