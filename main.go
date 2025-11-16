package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ShabarishRamaswamy/Srotas-Log-Aggregation-Simulator/src/methods"
	"github.com/ShabarishRamaswamy/Srotas-Log-Aggregation-Simulator/src/router"
)

var PORT string = ":8000"

func main() {
	var writerToPushInto io.Writer
	ctx, cancelCtx := context.WithCancel(context.WithValue(context.Background(), "writerToPushInto", writerToPushInto))

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, os.Interrupt)

	mainTicker := make(chan bool)
	go methods.TickEveryT(ctx, time.Second*5, mainTicker)

	fmt.Println("Server started on PORT: ", PORT)
	go func() {
		log.Fatal(http.ListenAndServe(PORT, router.GetNewRouter(ctx, mainTicker)))
	}()

	<-exit
	cancelCtx()
	fmt.Println("Graceful Shutdown")
	// Have a better implementation
	time.Sleep(time.Second)
}
