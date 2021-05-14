package usecase

import (
	"go.uber.org/zap"

	"github.com/ivantedja/xmarvel/characters"
	"github.com/ivantedja/xmarvel/marvels"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "characters-usecase")))
)

type Usecase struct {
	charactersCache characters.CacheRepository
	marvesUsecase   marvels.Usecase
}

func New(charactersCache characters.CacheRepository, marvesUsecase marvels.Usecase) characters.Usecase {
	return &Usecase{
		charactersCache: charactersCache,
		marvesUsecase:   marvesUsecase,
	}
}
