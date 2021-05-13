package usecase

import (
	"context"
	"time"

	"go.uber.org/zap"

	entity "github.com/ivantedja/xmarvel/entity"
)

type Filter struct {
	Offset        uint
	Limit         uint
	ModifiedSince time.Time
}

func (u *Usecase) Search(ctx context.Context, filter map[string]string) (*entity.CharacterCollection, error) {
	cc, err := u.marvelsRepository.Search(ctx, filter)
	if err != nil {
		logger.Error("search", zap.Error(err))
		return nil, err
	}
	return cc, nil
}
