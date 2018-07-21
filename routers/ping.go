package routers

import (
	"net/http"

	"github.com/json-iterator/go"
	"github.com/julienschmidt/httprouter"
)

var pong = struct {
	Status  string
	Details map[string]string
}{
	Status: "OK",
	Details: map[string]string{
		"DB": "TBD",
	},
}

func ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(&pong)
}
