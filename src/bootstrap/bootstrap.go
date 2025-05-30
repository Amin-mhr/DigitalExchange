package bootstrap

import (
	"DigitalExchange/src/api"
	"DigitalExchange/src/database"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Init() (err error) {

	defer func() {
		log.Println("Goodbye!")
		os.Exit(0)
	}()

	err = database.Init()
	if err != nil {
		log.Fatalf("Database Service: Failed to Initialize. %v", err)
	}
	log.Println("Database Service: Initialized Successfully.")

	defer func() {
		if err = database.GetInstance().Close(); err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}()

	// Initialize API
	go func() {
		err = api.Init()
		if err != nil {
			log.Fatalf("API Service: Failed to Initialize. %v", err)
		}
		log.Println("API Service: Initialized Successfully.")
	}()

	log.Println("Application is now running.\nPress CTRL-C to exit")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	log.Println("Application is shutting down...")

	return
}
