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
		tmpArrInt, err := u.asyncRetrieveCharacters(ctx, remaining, 1, 100)
		if err != nil {
			return nil, err
		}

		arrInt = append(arrInt, tmpArrInt...)
	}

	// Write list of character ID(s) to cache
	u.writeToCache(ctx, "marvels-characters", arrInt, 24*time.Hour)

	return arrInt, nil
}

//func (u *Usecase) syncRetrieveCharacters(ctx context.Context, remaining, iterationFrom, perBatch int) ([]uint, error) {
//	var arrInt []uint
//
//	for i := iterationFrom; i <= int(math.Ceil(float64(remaining)/float64(perBatch))); i++ {
//		cc, err := u.retrieveCharacters(ctx, fmt.Sprint(perBatch), fmt.Sprint(perBatch*i))
//		if err != nil {
//			return nil, err
//		}
//
//		for _, c := range cc.Data.Results {
//			arrInt = append(arrInt, c.ID)
//			// Write single character cache
//			u.writeToCache(ctx, "marvels-characters-"+fmt.Sprint(c.ID), c, 24*time.Hour)
//		}
//	}
//
//	return arrInt, nil
//}

func (u *Usecase) asyncRetrieveCharacters(ctx context.Context, remaining, iterationFrom, perBatch int) ([]uint, error) {
	var arrErr []error
	var arrInt []uint
	var waitgroup sync.WaitGroup

	queue := make(chan []entity.Character, 1)
	queueErr := make(chan error, 1)

	for i := iterationFrom; i <= int(math.Ceil(float64(remaining)/float64(perBatch))); i++ {
		waitgroup.Add(1)
		go func(i int) {
			cc, err := u.retrieveCharacters(ctx, fmt.Sprint(perBatch), fmt.Sprint(perBatch*i))
			if err != nil {
				queueErr <- err
				return
			}
			queue <- cc.Data.Results
		}(i)
	}

	go func() {
		for chrs := range queue {
			for _, c := range chrs {
				arrInt = append(arrInt, c.ID)
				// Write single character cache
				u.writeToCache(ctx, "marvels-characters-"+fmt.Sprint(c.ID), c, 24*time.Hour)
			}
			waitgroup.Done()
		}
	}()

	go func() {
		for qerr := range queueErr {
			arrErr = append(arrErr, qerr)
			waitgroup.Done()
		}
	}()

	waitgroup.Wait()

	// Just return any error
	if len(arrErr) > 0 {
		return nil, arrErr[0]
	}

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
