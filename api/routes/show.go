package routes

import (
	"fmt"
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

	paragraphs := post.Content.BodyModel.Paragraphs

	p := converters.ConvertParagraphs(paragraphs)

	desc := ""
	if len(paragraphs) >= 0 {
		for _, p := range paragraphs {
			if p.Type == "H3" || p.Type == "P" {
				desc = p.Text
				break
			}
		}
	}

	imgUrl := ""
	for _, p := range paragraphs {
		if p.Type == "IMG" {
			imgUrl = fmt.Sprintf("https://miro.medium.com/v2/resize:fit:1200/%s", p.Metadata.ID)
			break
		}
	}

	return c.Render("show", fiber.Map {
		"Title": post.Title,
		"UserId": post.Creator.ID,
		"Author": post.Creator.Name,
		"PublishDate": publishDate.Format(time.DateOnly),
		"Paragraphs": template.HTML(p),
		"Description": desc,
		"Path": c.Path(),
		"Image": imgUrl,
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