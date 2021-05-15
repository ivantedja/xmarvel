package usecase

import (
	"context"
	"encoding/json"
	"github.com/ivantedja/xmarvel/characters"
	"github.com/ivantedja/xmarvel/marvels"
	"time"
)

var (
//logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "characters-usecase")))
)

type Usecase struct {
	charactersCacheRepository characters.CacheRepository
	marvelsUsecase            marvels.Usecase
}

func New(charactersCache characters.CacheRepository, marvelsUsecase marvels.Usecase) characters.Usecase {
	return &Usecase{
		charactersCacheRepository: charactersCache,
		marvelsUsecase:            marvelsUsecase,
	}
}

func (u *Usecase) writeToCache(ctx context.Context, key string, val interface{}, expiration time.Duration) {
	arrJson, _ := json.Marshal(val)
	_ = u.charactersCacheRepository.Set(ctx, key, string(arrJson), expiration)
}
