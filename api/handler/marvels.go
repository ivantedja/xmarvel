package handler

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/go-chi/chi"

	"github.com/ivantedja/xmarvel/marvels"
)

type Marvels struct {
	*chi.Mux
	repository marvels.MarvelsRepository
	marvels    marvels.Usecase
}

func (m Marvels) Index(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	// TODO: adjust the filter accordingly
	filter := make(map[string]string)
	filter["limit"] = "2"
	filter["modifiedSince"] = "2015-04-28"

	cc, err := m.marvels.Search(ctx, filter)
	if err != nil {
		logger.Error("index", zap.Error(err))
		render(w, ErrBadRequest, 400)
		return
	}

	render(w, cc, 200)
}

func NewMarvels(repository marvels.MarvelsRepository, marvels marvels.Usecase) Marvels {
	h := Marvels{
		Mux:        chi.NewMux(),
		repository: repository,
		marvels:    marvels,
	}

	h.Get("/", h.Index)
	return h
}
