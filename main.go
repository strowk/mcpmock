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
	if len(os.Args) < 3 {
		log.Fatal("Usage: mcpmock serve <path to scenarios folder>")
	}
	if os.Args[1] != "serve" {
		log.Fatal("Usage: mcpmock serve <path to scenarios folder>")
	}
	srv := mockserver.NewMockServer(os.Args[2])
	c := make(chan os.Signal, 1)
	go func() {
		log.Println("mcpmock: Started server, press Ctrl+C to shutdown")
		<-c
		ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
		err := srv.Stop(ctx)
		if err != nil {
			log.Fatalf("mcpmock: Error stopping server: %v", err)
		}
	}()
	signal.Notify(c, os.Interrupt)
	err := srv.Run()
	if err != nil {
		log.Fatalf("mcpmock: Error running server: %v", err)
	}
}
