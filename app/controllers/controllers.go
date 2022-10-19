package controllers

import (
	"learn_db/pkg/utils"
	"learn_db/platform/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func SelectObject(variable *database.Variable, queryDb *string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		result, err := variable.Db.Exec(*queryDb)

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
			Data:       result,
		})
	}

}
