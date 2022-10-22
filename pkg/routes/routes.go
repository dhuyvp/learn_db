package routes

import (
	"learn_db/app/controllers"
	"learn_db/platform/database"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App, variable *database.Variable) {
	groupAPI := app.Group("/api")
	route := groupAPI.Group("/route")

	route.Post("/select", controllers.QuerySelect(variable))
	route.Post("/insert", controllers.QueryInsert(variable))

	route.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Successfull!")
	})
}
