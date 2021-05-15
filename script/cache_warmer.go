package main

import (
	"context"
	"github.com/joho/godotenv"
	"os"
	"time"

	"github.com/ivantedja/xmarvel/api"
	"github.com/ivantedja/xmarvel/characters"
	crepo "github.com/ivantedja/xmarvel/characters/repository"
	cusecase "github.com/ivantedja/xmarvel/characters/usecase"
	"github.com/ivantedja/xmarvel/lib"
	"github.com/ivantedja/xmarvel/marvels"
	mrepo "github.com/ivantedja/xmarvel/marvels/repository"
)

func main() {
	_ = godotenv.Load(".env")

	var (
		ctx               = context.Background()
		marvelsRepository = initMarvelsRepository()
		marvelsUsecase    = api.NewMarvels(marvelsRepository)
		cacheRepository   = initCacheRepository()
		charactersUsecase = cusecase.New(cacheRepository, marvelsUsecase)
	)
	charactersUsecase.Search(ctx)
}

func initMarvelsRepository() marvels.MarvelsRepository {
	baseUrl := os.Getenv("MARVEL_HOST")
	publicKey := os.Getenv("MARVEL_PUBLIC_KEY")
	privateKey := os.Getenv("MARVEL_PRIVATE_KEY")
	timeout := 5 * time.Second
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
