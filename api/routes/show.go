package routes

import "github.com/gofiber/fiber/v2"

func show(c *fiber.Ctx) error {
	return nil
}

func index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map {
		"Title": "medium.rip",
	})
}

func RegisterRoutes(app *fiber.App) {
	app.Get("/", index)
	app.Get("/:id", show)
}