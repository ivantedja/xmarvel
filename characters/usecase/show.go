package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ivantedja/xmarvel/entity"
)

func (u *Usecase) Show(ctx context.Context, ID int) (*entity.Character, error) {
	var chr entity.Character

	val, _ := u.charactersCacheRepository.Get(ctx, "marvels-characters-"+fmt.Sprint(ID))
	if val != "" {
		jerr := json.Unmarshal([]byte(val), &chr)
		if jerr != nil {
			return nil, jerr
		}
		return &chr, nil
	}

	c, err := u.marvelsUsecase.Show(ctx, ID)
	if err != nil {
		return nil, err
	}

	u.writeToCache(ctx, "marvels-characters-"+fmt.Sprint(c.ID), c, 24*time.Hour)

	return c, nil
}
