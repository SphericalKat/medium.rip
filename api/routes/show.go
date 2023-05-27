package routes

import (
	"fmt"
	"log"
	"strings"
	"time"

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

	post := e.Data.Post
	publishDate := time.UnixMilli(e.Data.Post.CreatedAt)
	log.Println(publishDate)

	var sb strings.Builder

	for _, node := range post.Content.BodyModel.Paragraphs {
		switch node.Type {
		case "H3":
			sb.WriteString(fmt.Sprintf("<h3>%s</h3>", node.Text))
		}
	}

	return c.Render("show", fiber.Map {
		"Title": post.Title,
		"UserId": post.Creator.ID,
		"Author": post.Creator.Name,
		"PublishDate": publishDate.Format(time.DateOnly),
		"Nodes": post.Content.BodyModel.Paragraphs,
	})
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