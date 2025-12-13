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
    //"github.com/gofiber/fiber/v2/middleware/proxy"
	//"crypto/tls"
	"github.com/gofiber/template/html/v2"
)

func RegisterRoutes(ctx context.Context, wg *sync.WaitGroup, engine *html.Engine, fs http.FileSystem) {
	// more Fiber options at https://docs.gofiber.io/api/fiber/
	app := fiber.New(fiber.Config{
		StreamRequestBody:     true,
		ServerHeader:          "Katbox",
		AppName:               "Katbox",
		DisableStartupMessage: false,
		Views: engine,
		Network: "tcp",
	})

	// static file server
	app.Use("/assets", filesystem.New(filesystem.Config{
		Root: fs,
		Browse: config.Conf.Env == "dev",
	}))

	/*
	if config.Conf.Proxy != "" {
		proxy.WithTlsConfig(&tls.Config{ InsecureSkipVerify: true,})
		app.Use(proxy.Balancer(proxy.Config{
			Servers: []string{ config.Conf.Proxy, },
			ModifyRequest: func(c *fiber.Ctx) error {
				//c.Request().Header.Add("X-Real-IP", c.IP())
				//return nil
				//req_code, req_body, req_errs := c.Request().String()
				//log.Printf("RESPONSE: %s | %s | %s ", req_code, req_body, req_errs)
				//log.Printf("REQUEST: %s", c.Request().Body())
				log.Printf("RESP: %s", string(c.Request().String()))
				return nil
			},
			ModifyResponse: func(c *fiber.Ctx) error {
				//c.Response().Header.Del(fiber.HeaderServer)
				//return nil
				//res_code, res_body, res_errs := c.Response().String()
				//log.Printf("RESPONSE: %s | %s | %s ", res_code, res_body, res_errs)
				log.Printf("RESP: %s", string(c.Response().String()))
				return nil
			},
		}))
		log.Printf("Using proxy: %s", config.Conf.Proxy)
	}
	*/
	
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
