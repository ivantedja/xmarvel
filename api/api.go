package api

import (
	"github.com/go-chi/chi"
	chimid "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/ivantedja/xmarvel/api/handler"
	"github.com/ivantedja/xmarvel/characters"
	cusecase "github.com/ivantedja/xmarvel/characters/usecase"
	"github.com/ivantedja/xmarvel/marvels"
	musecase "github.com/ivantedja/xmarvel/marvels/usecase"
)

func NewMarvels(repository marvels.MarvelsRepository) marvels.Usecase {
	return musecase.New(repository)
}

func NewMux(cacheRepository characters.CacheRepository, musecase marvels.Usecase) *chi.Mux {
	var (
		mux               = chi.NewMux()
		charactersUsecase = cusecase.New(cacheRepository, musecase)
		marvelsHandler    = handler.NewMarvels(cacheRepository, charactersUsecase)
	)

	mux.Use(chimid.RequestID)
	mux.Use(chimid.RealIP)
	mux.Use(chimid.Recoverer)
	mux.Use(cors.AllowAll().Handler)

	mux.Mount("/characters", marvelsHandler)

	return mux
}
