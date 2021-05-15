package handler

import (
	"net/http"
	"strconv"

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

func (c Characters) Show(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		qID = chi.URLParam(r, "ID")
	)

	ID, err := strconv.Atoi(qID)
	if err != nil {
		logger.Error("show", zap.Error(err))
		render(w, ErrBadRequest, 400)
		return
	}

	cc, err := c.characters.Show(ctx, ID)
	if err != nil {
		logger.Error("show", zap.Error(err))
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
	h.Get("/{ID}", h.Show)
	return h
}
