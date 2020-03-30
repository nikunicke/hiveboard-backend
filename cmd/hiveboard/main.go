package main

import (
	"log"

	"github.com/nikunicke/hiveboard/bolt"
	"github.com/nikunicke/hiveboard/http"
)

func main() {
	// http.Run()
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	httpServer := http.NewServer()
	httpServer.Addr = ":3000"
	httpServer.EventService = bolt.NewEventService()
	err := httpServer.Open()
	if err != nil {
		return err
	}
	return nil
}
