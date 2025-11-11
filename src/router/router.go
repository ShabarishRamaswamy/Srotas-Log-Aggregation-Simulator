package router

import (
	"net/http"

	"github.com/ShabarishRamaswamy/Srotas-Log-Aggregation-Simulator/src/methods"
)

func greetHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("\u0939\u0947llo Worl\u0921"))
}

func getALog(w http.ResponseWriter, r *http.Request) {
	methods.GetALog(w)
}

func GetNewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/hello", greetHello)
	router.HandleFunc("/getALog", getALog)

	return router
}
