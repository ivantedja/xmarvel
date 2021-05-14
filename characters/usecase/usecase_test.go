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

func initiateCharacterCollectionResponse() *entity.CharacterCollection {
	return &entity.CharacterCollection{
		Code:   200,
		Status: "Ok",
		Data:   entity.CharacterData{},
	}
}

//func initiateCharactersResponse() []*entity.Character {
//	return []*entity.Character{
//		{
//			ID:          1011198,
//			Name:        "Agents of Atlas",
//			Description: "See my map",
//		},
//		{
//			ID:          1010801,
//			Name:        "Ant-Man (Scott Lang)",
//			Description: "Build colony",
//		},
//	}
//}
