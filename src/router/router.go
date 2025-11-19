package router

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ShabarishRamaswamy/Srotas-Log-Aggregation-Simulator/src/methods"
	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
)

var MainTicker chan bool

func greetHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("\u0939\u0947llo Worl\u0921"))
}

func handleSocketListen(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer c.CloseNow()

	var v any
	err = wsjson.Read(context.Background(), c, &v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	log.Printf("Received: %+v", v)
	for <-MainTicker {
		wsjson.Write(context.Background(), c, "Well, Hello There!")
	}

	wsjson.Write(context.Background(), c, v)

	c.Close(websocket.StatusNormalClosure, "")
}

func GetNewRouter(ctx context.Context, ticker chan bool) *http.ServeMux {
	MainTicker = ticker
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

	router.HandleFunc("/listen", handleSocketListen)

	return router
}
