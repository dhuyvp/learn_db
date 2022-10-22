package controllers

import (
	"fmt"
	"learn_db/app/models"
	"learn_db/pkg/utils"
	"learn_db/platform/database"
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

func QuerySelect(variable *database.Variable) fiber.Handler {
	queryDb := "SELECT * FROM Person WHERE id < 15"

	return func(c *fiber.Ctx) error {
		queryResult := []models.Person{}
		err := variable.Db.Select(&queryResult, queryDb)

		if err != nil {
			log.Println("Error to select database", err)
			return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		return c.Status(fiber.StatusOK).JSON(&utils.Response{
			Success:    true,
			StatusCode: fiber.StatusOK,
			Data:       queryResult,
		})
	}

}

func QueryInsert(variable *database.Variable) fiber.Handler {
	rand.Seed(time.Now().Unix())

	fName := utils.RandString(4)
	lName := utils.RandString(6)
	randAge := utils.RandInt(2)
	queryDb := fmt.Sprintf("INSERT INTO Person (FirstName, LastName, Age) VALUES ('%s', '%s', '%s')", fName, lName, randAge)

	fmt.Println("Query db: ", queryDb)
	return func(c *fiber.Ctx) error {
		result, errInsert := variable.Db.Exec(queryDb)
		if errInsert != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		resId, errId := result.LastInsertId()
		if errId != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		queryResult := []models.Person{}
		err := variable.Db.Select(&queryResult, "Select * from Person Where id = ?;", resId)

		if err != nil {
			log.Println("Error when inserting data!", err)
			return c.Status(fiber.StatusBadRequest).JSON(&utils.Response{
				Success:    false,
				StatusCode: fiber.StatusBadRequest,
			})
		}

		return c.Status(fiber.StatusOK).JSON(&utils.Response{
			Success:    true,
			StatusCode: fiber.StatusOK,
			Data:       queryResult,
		})
	}

}
