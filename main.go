package main

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/medium.rip/api"
	"github.com/medium.rip/internal/config"
	"github.com/medium.rip/internal/lifecycle"
)

func main() {
	// load config
	config.Load()

	// cancellable context
	wg := sync.WaitGroup{}
	ctx, cancelFunc := context.WithCancel(context.Background())

	wg.Add(1)
	go api.RegisterRoutes(ctx, &wg)

	wg.Add(1)
	go lifecycle.ShutdownListener(&wg, &cancelFunc)

	wg.Wait()

	log.Info("Graceful shutdown complete")
}
