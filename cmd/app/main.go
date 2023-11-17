package main

import (
	"currency-converter-rev2/server"
	"log"
)

func main() {
	app, err := server.NewApp()
	if err != nil {
		log.Fatalf("Failed to initialize the application: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Printf("Failed to run the application: %v", err)
	}
}
