package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/medium.rip/pkg/client"
)

func show(c *fiber.Ctx) error {
	postId := c.Params("id", "")
	if postId == "" {
		return c.Redirect("/")
	}

	e, err := client.PostData(postId)
	if err != nil {
		return err
	}

	return c.JSON(e)
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