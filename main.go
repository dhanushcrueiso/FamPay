package main

import (
	"FAMPAY/config"
	"FAMPAY/internal/database/db"
	"FAMPAY/internal/routes"
	"FAMPAY/internal/services"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber"
)

func main() {
	//setting env to dev as we need to fetch data from dev.json
	env := "dev"

	var file *os.File
	var err error

	file, err = os.Open(env + ".json")
	if err != nil {
		log.Println("Unable to open file. Err:", err)
		os.Exit(1)
	}
	//parsing json with the config and passing the dev.json values from here
	var cnf *config.Config
	config.ParseJSON(file, &cnf)
	config.Set(cnf)

	db.Init(&db.Config{
		URL:       cnf.DatabaseURL,
		MaxDBConn: cnf.MaxDBConn,
	})

	app := fiber.New()

	routes.SetupRoutes(app)
	fmt.Printf("Server is running on port %s\n", cnf.Port)

	go func() {
		if err := app.Listen(cnf.Port); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// This will prevent the main function from exiting immediately

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			go services.FetchAndStoreVideos()
		}
	}
}
