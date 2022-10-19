package routes

import (
	"learn_db/app/controllers"
	"learn_db/platform/database"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App, variable *database.Variable) {
	groupAPI := app.Group("/group/")
	route := groupAPI.Group("/route")

	querySelect := "SELECT * FROM Cities;"
	route.Post("", controllers.SelectObject(variable, &querySelect))

	route.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Succesfull!")
	})
}
