package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ShabarishRamaswamy/Srotas-Log-Aggregation-Simulator/src/methods"
)

func greetHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("\u0939\u0947llo Worl\u0921"))
}

func GetNewRouter(ctx context.Context, ticker chan bool) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/hello", greetHello)
	router.HandleFunc("/getALog", func(w http.ResponseWriter, r *http.Request) {
		i := 0
		for <-ticker {
			if i == 3 {
				return
			}
			fmt.Println("Ticked again", i)
			w.Write([]byte("Hey"))
			methods.SendALog(w)
			i += 1
		}
	})

	return router
}
