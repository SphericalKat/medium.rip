package routes

import (
	"html/template"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/medium.rip/pkg/client"
	"github.com/medium.rip/pkg/converters"
)

func show(c *fiber.Ctx) error {
	postId := converters.ConvertId(c)
	if postId == "" {
		return c.Status(422).SendString("Invalid post ID")
	}

	e, err := client.PostData(postId)
	if err != nil {
		return err
	}

	post := e.Data.Post
	publishDate := time.UnixMilli(e.Data.Post.CreatedAt)

	p := converters.ConvertParagraphs(post.Content.BodyModel.Paragraphs)

	return c.Render("show", fiber.Map {
		"Title": post.Title,
		"UserId": post.Creator.ID,
		"Author": post.Creator.Name,
		"PublishDate": publishDate.Format(time.DateOnly),
		"Paragraphs": template.HTML(p),
	})
}

func index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map {
		"Title": "medium.rip",
	})
}

func RegisterRoutes(app *fiber.App) {
	app.Get("/", index)
	app.Get("/*", show)
}