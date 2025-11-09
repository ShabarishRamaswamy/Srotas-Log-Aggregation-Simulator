package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ShabarishRamaswamy/Srotas-Log-Aggregation-Simulator/src/router"
)

var PORT string = ":8000"

func main() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, os.Interrupt)

	fmt.Println("Server started on PORT: ", PORT)
	go func() {
		log.Fatal(http.ListenAndServe(PORT, router.GetNewRouter()))
	}()

	<-exit
	fmt.Println("Graceful Shutdown")
}
