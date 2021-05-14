package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func (u *Usecase) Search(ctx context.Context) ([]uint, error) {
	var arrInt []uint

	val, _ := u.charactersCacheRepository.Get(ctx, "marvels-characters")
	if val != "" {
		jerr := json.Unmarshal([]byte(val), &arrInt)
		if jerr != nil {
			return nil, jerr
		}
		return arrInt, nil
	}

	filter := make(map[string]string)
	filter["limit"] = "100"

	cc, err := u.marvelsUsecase.Search(ctx, filter)
	if err != nil {
		return nil, err
	}

	for _, c := range cc.Data.Results {
		arrInt = append(arrInt, c.ID)
		u.writeToCache(ctx, "marvels-characters-"+fmt.Sprint(c.ID), c, 24*time.Hour)
	}

	u.writeToCache(ctx, "marvels-characters", arrInt, 24*time.Hour)

	return arrInt, nil
}

func (u *Usecase) writeToCache(ctx context.Context, key string, val interface{}, expiration time.Duration) {
	arrJson, _ := json.Marshal(val)
	_ = u.charactersCacheRepository.Set(ctx, key, string(arrJson), expiration)
}
