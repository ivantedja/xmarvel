package main

import (
	"context"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/ivantedja/xmarvel/api"
	"github.com/ivantedja/xmarvel/marvels"
	"github.com/ivantedja/xmarvel/marvels/repository"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "main")))
	shutdowns []func() error
)

func main() {
	var (
		ctx               = context.Background()
		port              = os.Getenv("PORT")
		marvelsRepository = initMarvelsRepository()
		mux               = api.NewMux(marvelsRepository)
		server            = http.Server{
			Addr:    ":" + port,
			Handler: mux,
		}
		shutdown = make(chan struct{})
	)

	go gracefulShutdown(ctx, &server, shutdown)

	logger.Info("server starting: http://localhost" + server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatal("server error", zap.Error(err))
	}

	<-shutdown
}

func initMarvelsRepository() marvels.MarvelsRepository {
	baseUrl := os.Getenv("MARVEL_HOST")
	publicKey := os.Getenv("MARVEL_PUBLIC_KEY")
	privateKey := os.Getenv("MARVEL_PRIVATE_KEY")
	timeout := 5 * time.Second
	repository := repository.NewAPI(baseUrl, publicKey, privateKey, timeout)
	return repository
}

func gracefulShutdown(ctx context.Context, server *http.Server, shutdown chan struct{}) {
	var (
		sigint = make(chan os.Signal, 1)
	)

	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	<-sigint

	logger.Info("shutting down server gracefully")

	// stop receiving any request.
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("shutdown error", zap.Error(err))
	}

	// close any other modules.
	for i := range shutdowns {
		_ = shutdowns[i]()
	}

	close(shutdown)
}
