package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/ivantedja/xmarvel/entity"
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

	// Run first batch (100 records) to determine the total
	cc, err := u.retrieveCharacters(ctx, "100", "0")
	if err != nil {
		return nil, err
	}

	for _, c := range cc.Data.Results {
		arrInt = append(arrInt, c.ID)
		// Write single character cache
		u.writeToCache(ctx, "marvels-characters-"+fmt.Sprint(c.ID), c, 24*time.Hour)
	}

	// Calculate remaining batches
	var total, remaining int
	total = int(cc.Data.Total)
	remaining = total - 100

	// Run remaining batches if remaining exceed initial limit (100)
	if remaining > 100 {
		var waitgroup sync.WaitGroup
		waitgroup.Add(remaining)

		queue := make(chan uint, 1)

		for i := 1; i <= int(math.Ceil(float64(remaining)/100)); i++ {
			go func(i int) {
				cc, err := u.retrieveCharacters(ctx, "100", fmt.Sprint(100*i))
				if err != nil {
					return
				}

				for _, c := range cc.Data.Results {
					queue <- c.ID
					// Write single character cache
					u.writeToCache(ctx, "marvels-characters-"+fmt.Sprint(c.ID), c, 24*time.Hour)
				}
			}(i)
		}

		go func() {
			for i := range queue {
				arrInt = append(arrInt, i)
				waitgroup.Done()
			}
		}()

		waitgroup.Wait()
	}

	// Write list of character ID(s) to cache
	u.writeToCache(ctx, "marvels-characters", arrInt, 24*time.Hour)

	return arrInt, nil
}

func (u *Usecase) retrieveCharacters(ctx context.Context, limit, offset string) (*entity.CharacterCollection, error) {
	filter := make(map[string]string)
	filter["limit"] = limit
	filter["offset"] = offset

	cc, err := u.marvelsUsecase.Search(ctx, filter)
	if err != nil {
		return nil, err
	}

	return cc, nil
}

func (u *Usecase) writeToCache(ctx context.Context, key string, val interface{}, expiration time.Duration) {
	arrJson, _ := json.Marshal(val)
	_ = u.charactersCacheRepository.Set(ctx, key, string(arrJson), expiration)
}
