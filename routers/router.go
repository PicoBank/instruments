package routers

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var router httprouter.Router

func Start() {
	router := httprouter.New()

	router.GET("/", index)
	router.GET("/ping", ping)

	log.Fatal(http.ListenAndServe(":8080", router))

}
