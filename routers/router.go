package routers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/json-iterator/go"
	"github.com/julienschmidt/httprouter"
)

type pong struct {
	Status  string
	Details map[string]string
}

var pongValue = pong{
	Status: "OK",
	Details: map[string]string{
		"DB":     "OK",
		"Cache":  "OK",
		"Coffee": "OK",
	},
}

var router httprouter.Router

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Securities Master File 3.14!\n")
}

func ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&pongValue)
}

func Start() {
	router := httprouter.New()

	router.GET("/", index)
	router.GET("/ping", ping)

	log.Fatal(http.ListenAndServe(":8080", router))

}
