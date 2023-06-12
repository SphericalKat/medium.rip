package main

import (
	"context"
	"embed"
	"io/fs"
	"net/http"
	"sync"

	"github.com/gofiber/template/html/v2"
	log "github.com/sirupsen/logrus"

	"github.com/medium.rip/api"
	"github.com/medium.rip/internal/config"
	"github.com/medium.rip/internal/lifecycle"
)

//go:embed frontend
var frontend embed.FS

//go:embed frontend/dist/assets
var static embed.FS

func main() {
	// load config
	config.Load()

	// create template engine
	tmplFs, err := fs.Sub(frontend, "frontend/dist")
	if err != nil {
		log.Fatalf("error loading template: %v\n", err)
	}

	engine := html.NewFileSystem(http.FS(tmplFs), ".html")

	// create static file server
	staticFs, err := fs.Sub(static, "frontend/dist/assets")
	if err != nil {
		log.Fatalf("error loading static assets: %v\n", err)
	}

	staticHttp := http.FS(staticFs)

	// cancellable context
	wg := sync.WaitGroup{}
	ctx, cancelFunc := context.WithCancel(context.Background())

	wg.Add(1)
	go api.RegisterRoutes(ctx, &wg, engine, staticHttp)

	wg.Add(1)
	go lifecycle.ShutdownListener(&wg, &cancelFunc)

	wg.Wait()

	log.Info("Graceful shutdown complete")
}
