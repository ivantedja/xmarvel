package usecase

import (
	"context"
	"encoding/json"
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
		return arrInt, nil
	}

	for _, c := range cc.Data.Results {
		arrInt = append(arrInt, c.ID)
	}

	arrJson, _ := json.Marshal(arrInt)
	u.charactersCacheRepository.Set(ctx, "marvels-characters", string(arrJson), 24*time.Hour)

	return arrInt, nil
}
