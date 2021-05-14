package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"

	"github.com/ivantedja/xmarvel/characters"
)

type Marvels struct {
	*chi.Mux
	cacheRepository characters.CacheRepository
	characters      characters.Usecase
}

func (m Marvels) Index(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	cc, err := m.characters.Search(ctx)
	if err != nil {
		logger.Error("index", zap.Error(err))
		render(w, ErrBadRequest, 400)
		return
	}

	render(w, cc, 200)
}

func NewMarvels(cacheRepository characters.CacheRepository, characters characters.Usecase) Marvels {
	h := Marvels{
		Mux:             chi.NewMux(),
		cacheRepository: cacheRepository,
		characters:      characters,
	}

	h.Get("/", h.Index)
	return h
}
