package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"

	"github.com/betosardinha/antixy-api/internal/adapters/database"
	"github.com/betosardinha/antixy-api/internal/app"
)

func main() {
	database, err := initDatabase()
	if err != nil {
		log.Fatal(err)
	}

	application := app.New(database)

	if err := runServer(application); err != nil {
		log.Fatal(err)
	}
}

func initDatabase() (*gorm.DB, error) {
	connection, err := database.Connect()
	if err != nil {
		return nil, fmt.Errorf("connect database: %w", err)
	}

	log.Println("Database connected")

	return connection, nil
}

func runServer(application *app.App) error {
	log.Println("Server running on :8080")

	if err := application.Router.Run(":8080"); err != nil {
		return fmt.Errorf("run server: %w", err)
	}

	return nil
}
