package main

import (
	"fmt"
	"learn_db/pkg/routes"
	"learn_db/platform/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error to loading .env file")
	}

	variable, err := database.New()

	if err != nil {
		log.Fatal(fmt.Println(err))
	} else {
		fmt.Println("Database connected!")
	}

	app := fiber.New()

	routes.PublicRoutes(app, variable)

	if err := app.Listen(":1206"); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

}
