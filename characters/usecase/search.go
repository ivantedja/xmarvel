package usecase

import (
	"context"
	"encoding/json"
)

func (u *Usecase) Search(ctx context.Context) ([]uint, error) {
	val, _ := u.charactersCacheRepository.Get(ctx, "marvels-characters")
	if val != "" {
		var arrInt []uint
		jerr := json.Unmarshal([]byte(val), &arrInt)
		if jerr != nil {
			return nil, jerr
		}
		return arrInt, nil
	}

	return []uint{}, nil
}
