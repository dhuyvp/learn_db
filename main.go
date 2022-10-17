package main

import (
	"fmt"
	"learn_db/app/models"
	"learn_db/platform/database"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error to loading .env file")
	}

	variable, err := database.New()

	if err != nil {
		fmt.Println(err.Error())
	}

	interf := models.Person{
		FirstName: "Huy",
		LastName:  "Nguyen",
		Age:       20,
	}

	variable.InsertObject("Person", interf, false)

	// var s string
	//s = "\"string\""
	s := "string"
	log.Fatal(fmt.Println(s))

}
