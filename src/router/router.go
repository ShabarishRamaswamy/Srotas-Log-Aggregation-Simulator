package router

import (
	"net/http"
)

func greetHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("\u0939\u0947llo Worl\u0921"))
}

func GetNewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/hello", greetHello)

	return router
}
