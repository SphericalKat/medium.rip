package api

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/medium.rip/api/routes"
	"github.com/medium.rip/internal/config"
	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html/v2"
)

func RegisterRoutes(ctx context.Context, wg *sync.WaitGroup, engine *html.Engine, fs http.FileSystem) {
	app := fiber.New(fiber.Config{
		StreamRequestBody:     true,
		ServerHeader:          "Katbox",
		AppName:               "Katbox",
		DisableStartupMessage: true,
		Views: engine,
	})

	// static file server
	app.Use("/assets", filesystem.New(filesystem.Config{
		Root: fs,
		Browse: config.Conf.Env == "dev",
	}))

	routes.RegisterRoutes(app)

	go func(app *fiber.App) {
		log.Printf("Starting http server at: http://localhost:%s", config.Conf.Port)
		if err := app.Listen(fmt.Sprintf(":%s", config.Conf.Port)); err != nil {
			log.Fatalf("Unable to start http server: %s", err)
		}
	}(app)

	// listen for context cancellation
	<-ctx.Done()

	// shut down http server
	log.Info("Gracefully shutting down http server...")
	if err := app.Shutdown(); err != nil {
		log.Warn("Server shutdown Failed: ", err)
	}
	wg.Done()
}