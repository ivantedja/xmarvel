package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"go.uber.org/zap"

	"github.com/ivantedja/xmarvel/characters"
	"github.com/ivantedja/xmarvel/entity"
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
		render(w, entity.ErrBadRequest{Message: "Bad Request"}, 400)
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
		render(w, entity.ErrBadRequest{Message: "Bad Request"}, 400)
		return
	}

	cc, err := c.characters.Show(ctx, ID)
	if errors.Is(err, entity.ErrNotFound{Message: "Not Found"}) {
		logger.Error("show", zap.Error(err))
		render(w, err, 404)
		return
	}

	if err != nil {
		logger.Error("show", zap.Error(err))
		render(w, entity.ErrBadRequest{Message: "Bad Request"}, 400)
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
