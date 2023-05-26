package api

import "github.com/gofiber/fiber/v2"

func RegisterRoutes() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Listen(":3000")
}