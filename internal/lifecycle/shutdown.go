package lifecycle

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	log "github.com/sirupsen/logrus"
)

// ShutdownListener listens for shutdown OS signals and then gracefully stops tasks using the context
func ShutdownListener(
	wg *sync.WaitGroup,
	cf *context.CancelFunc,
) {
	// create channel to notify on system signals
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// wait for signal
	sig := <-termChan
	log.Printf("Received signal %v, gracefully shutting down services", sig)

	// close channel
	close(termChan)

	// cancel background context
	(*cf)()

	// mark shutdown task as done
	wg.Done()
}