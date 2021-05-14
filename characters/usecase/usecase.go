package usecase

import (
	"github.com/ivantedja/xmarvel/characters"
	"github.com/ivantedja/xmarvel/marvels"
)

var (
	//logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "characters-usecase")))
)

type Usecase struct {
	charactersCacheRepository characters.CacheRepository
	marvelsUsecase             marvels.Usecase
}

func New(charactersCache characters.CacheRepository, marvelsUsecase marvels.Usecase) characters.Usecase {
	return &Usecase{
		charactersCacheRepository: charactersCache,
		marvelsUsecase:             marvelsUsecase,
	}
}
