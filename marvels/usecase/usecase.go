package usecase

import (
	"go.uber.org/zap"

	"github.com/ivantedja/xmarvel/marvels"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "usecase")))
)

type Usecase struct {
	marvelsRepository marvels.MarvelsRepository
}

func New(marvelsRepository marvels.MarvelsRepository) marvels.Usecase {
	return &Usecase{
		marvelsRepository: marvelsRepository,
	}
}
