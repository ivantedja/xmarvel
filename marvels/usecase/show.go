package usecase

import (
	"context"

	"go.uber.org/zap"

	entity "github.com/ivantedja/xmarvel/entity"
)

func (u *Usecase) Show(ctx context.Context, ID int) (*entity.Character, error) {
	cc, err := u.marvelsRepository.Show(ctx, ID)
	if err != nil {
		logger.Error("show", zap.Error(err))
		return nil, err
	}
	return cc, nil
}
