package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"github.com/ivantedja/xmarvel/api"
	"github.com/ivantedja/xmarvel/characters"
	crepo "github.com/ivantedja/xmarvel/characters/repository"
	"github.com/ivantedja/xmarvel/lib"
	"github.com/ivantedja/xmarvel/marvels"
	mrepo "github.com/ivantedja/xmarvel/marvels/repository"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "main")))
	shutdowns []func() error
)

func main() {
	_ = godotenv.Load(".env")

	var (
		ctx               = context.Background()
		port              = os.Getenv("SERVICE_PORT")
		marvelsRepository = initMarvelsRepository()
		marvelsUsecase    = api.NewMarvels(marvelsRepository)
		cacheRepository   = initCacheRepository()

		mux    = api.NewMux(cacheRepository, marvelsUsecase)
		server = http.Server{
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
	timeout := 10 * time.Second
	repository := mrepo.NewAPI(baseUrl, publicKey, privateKey, timeout)
	return repository
}

func initCacheRepository() characters.CacheRepository {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	r := lib.NewRedis(redisHost, redisPort)
	repository := crepo.NewCache(r)
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
