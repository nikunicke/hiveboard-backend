package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/nikunicke/hiveboard"

	"github.com/nikunicke/hiveboard/api42"
	"github.com/nikunicke/hiveboard/http"
	"github.com/nikunicke/hiveboard/mongodb"
)

func main() {
	// Add configuration stuff here, i.e params and flags
	// or input from config files
	//testing mongo package
	// db := mongodb.NewMongoDB()
	// if err := db.Open("hiveboard"); err != nil {
	// 	log.Fatal(err)
	// }
	// if err := db.CheckConnection(); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Connected to DB")

	// if err := db.PostTest("events"); err != nil {
	// 	log.Fatal(err)
	// }

	// GET ALL TEST
	// db.FindAll("events")
	// ////////////////////////////////////////////

	if err := run(); err != nil {
		log.Fatal(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println(" --> Program shutting down")
}

func run() error {
	httpServer := http.NewServer()
	httpServer.Addr = ":3000"

	db := mongodb.NewMongoDB()
	if err := db.Open("hiveboard"); err != nil {
		log.Fatal(err)
	}

	// db.PostTest("events")

	httpServer.ES = *hiveboard.NewE()
	httpServer.ES.API42 = api42.NewEventService()
	httpServer.ES.Mongodb = mongodb.NewEventService(db)

	// httpServer.EventService = api42.NewEventService()
	httpServer.UserService = api42.NewUserService()
	err := httpServer.Open()
	if err != nil {
		return err
	}
	u := httpServer.URL()
	fmt.Printf("Server running on: %s\n", u.String())
	return nil
}
