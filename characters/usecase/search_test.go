package usecase_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/ivantedja/xmarvel/characters"
)

type SearchTestSuite struct {
	suite.Suite
	usecase characters.Usecase
	BaseTestSuite
}

func TestSearch(t *testing.T) {
	suite.Run(t, new(SearchTestSuite))
}

func (s *SearchTestSuite) SetupTest() {
	s.BaseTestSuite = getBaseTestSuite()
	s.usecase = initiateUsecase(&s.BaseTestSuite)
}

func (s *SearchTestSuite) TestCacheHitSuccess() {
	s.CacheRepository.On("Get", mock.Anything, "marvels-characters").Return(`[1011198, 1010801]`, nil).Once()
	s.MarvelsUsecase.On("Search", mock.Anything, map[string]string{}).Return(initiateCharacterCollectionResponse(), nil).Once()
	arr, merr := s.usecase.Search(context.Background())
	s.Assert().Equal(arr, []uint{1011198, 1010801})
	s.Assert().Nil(merr)
}

func (s *SearchTestSuite) TestCacheHitMarshalError() {
	s.CacheRepository.On("Get", mock.Anything, "marvels-characters").Return("zzz", nil).Once()
	s.MarvelsUsecase.On("Search", mock.Anything, map[string]string{}).Return(initiateCharacterCollectionResponse(), nil).Once()
	_, merr := s.usecase.Search(context.Background())
	s.Assert().NotNil(merr)
}

func (s *SearchTestSuite) TestCacheHitEmpty() {
	s.CacheRepository.On("Get", mock.Anything, "marvels-characters").Return("", nil).Once()
	s.MarvelsUsecase.On("Search", mock.Anything, map[string]string{}).Return(initiateCharacterCollectionResponse(), nil).Once()
	arr, merr := s.usecase.Search(context.Background())
	s.Assert().Equal(arr, []uint{})
	s.Assert().Nil(merr)
}
