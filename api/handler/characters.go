package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"

	"github.com/ivantedja/xmarvel/characters"
)

type Characters struct {
	*chi.Mux
	cacheRepository characters.CacheRepository
	characters      characters.Usecase
}

func (c Characters) Index(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	cc, err := c.characters.Search(ctx)
	if err != nil {
		logger.Error("index", zap.Error(err))
		render(w, ErrBadRequest, 400)
		return
	}

	render(w, cc, 200)
}

func NewCharacters(cacheRepository characters.CacheRepository, characters characters.Usecase) Characters {
	h := Characters{
		Mux:             chi.NewMux(),
		cacheRepository: cacheRepository,
		characters:      characters,
	}

	h.Get("/", h.Index)
	return h
}
