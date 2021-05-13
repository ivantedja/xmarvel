package api

import (
	"github.com/go-chi/chi"
	chimid "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/ivantedja/xmarvel/api/handler"
	"github.com/ivantedja/xmarvel/marvels"
	"github.com/ivantedja/xmarvel/marvels/usecase"
)

func NewMux(repository marvels.MarvelsRepository) *chi.Mux {
	var (
		mux            = chi.NewMux()
		marvels        = usecase.New(repository)
		marvelsHandler = handler.NewMarvels(repository, marvels)
	)

	mux.Use(chimid.RequestID)
	mux.Use(chimid.RealIP)
	mux.Use(chimid.Recoverer)
	mux.Use(cors.AllowAll().Handler)

	mux.Mount("/characters", marvelsHandler)

	return mux
}
