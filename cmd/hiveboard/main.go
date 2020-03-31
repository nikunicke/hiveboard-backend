package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/nikunicke/hiveboard/api42"
	"github.com/nikunicke/hiveboard/http"
)

func main() {
	// Add configuration stuff here, i.e params and flags
	// or input from config files

	if err := Run(); err != nil {
		log.Fatal(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println(" --> Program shutting down")
}

func Run() error {
	httpServer := http.NewServer()
	httpServer.Addr = ":3000"
	httpServer.EventService = api42.NewEventService()
	httpServer.UserService = api42.NewUserService()
	err := httpServer.Open()
	if err != nil {
		return err
	}
	return nil
}
