package usecase_test

import (
	"github.com/ivantedja/xmarvel/characters"
	cmocks "github.com/ivantedja/xmarvel/characters/mocks"
	"github.com/ivantedja/xmarvel/characters/usecase"
	"github.com/ivantedja/xmarvel/entity"
	mmocks "github.com/ivantedja/xmarvel/marvels/mocks"
)

type BaseTestSuite struct {
	CacheRepository *cmocks.CacheRepository
	MarvelsUsecase  *mmocks.Usecase
}

func getBaseTestSuite() BaseTestSuite {
	return BaseTestSuite{
		CacheRepository: new(cmocks.CacheRepository),
		MarvelsUsecase:  new(mmocks.Usecase),
	}
}

func initiateUsecase(base *BaseTestSuite) characters.Usecase {
	return usecase.New(
		base.CacheRepository,
		base.MarvelsUsecase,
	)
}

func initiateCharacterCollectionResponse(total, chars, offset uint) *entity.CharacterCollection {
	return &entity.CharacterCollection{
		Code:   200,
		Status: "Ok",
		Data: entity.CharacterData{
			Total:   total,
			Results: initiateCharactersResponse(chars, offset),
		},
	}
}

func initiateCharactersResponse(chars, offset uint) []entity.Character {
	var c []entity.Character
	var i uint
	for i = 1; i <= chars; i++ {
		c = append(c, entity.Character{
			ID:          1011000 + offset + i,
			Name:        "",
			Description: "",
		},
		)
	}
	return c
}
