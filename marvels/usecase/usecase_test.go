package usecase_test

import (
	"github.com/ivantedja/xmarvel/entity"
	"github.com/ivantedja/xmarvel/marvels"
	"github.com/ivantedja/xmarvel/marvels/mocks"
	"github.com/ivantedja/xmarvel/marvels/usecase"
)

type BaseTestSuite struct {
	MarvelsRepository *mocks.MarvelsRepository
}

func getBaseTestSuite() BaseTestSuite {
	return BaseTestSuite{
		MarvelsRepository: new(mocks.MarvelsRepository),
	}
}

func initiateUsecase(base *BaseTestSuite) marvels.Usecase {
	return usecase.New(
		base.MarvelsRepository,
	)
}

func initiateCharacterCollectionResponse() *entity.CharacterCollection {
	return &entity.CharacterCollection{
		Code:   200,
		Status: "Ok",
		Data:   entity.CharacterData{},
	}
}
