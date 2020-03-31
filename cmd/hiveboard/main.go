package main

import (
	"log"

	"github.com/nikunicke/hiveboard/bolt"
	"github.com/nikunicke/hiveboard/http"
)

func main() {
	// Add configuration stuff here, i.e params and flags
	// or input from config files

	if err := Run(); err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	httpServer := http.NewServer()
	httpServer.Addr = ":3000"
	httpServer.EventService = bolt.NewEventService()
	httpServer.UserService = bolt.NewUserService()
	err := httpServer.Open()
	if err != nil {
		return err
	}
	return nil
}
