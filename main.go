package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	mockserver "github.com/strowk/mcpmock/pkg/mockserver"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("mcpmock: Please provide a path to the folder with scenarios")
	}
	srv := mockserver.NewMockServer(os.Args[1])
	c := make(chan os.Signal, 1)
	go func() {
		log.Println("mcpmock: Started server, press Ctrl+C to shutdown")
		<-c
		ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
		srv.Stop(ctx)
	}()
	signal.Notify(c, os.Interrupt)
	srv.Run()
}
