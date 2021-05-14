package usecase_test

import (
	"context"
	"errors"
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
	s.CacheRepository.On("Get", mock.Anything, "marvels-characters").Return(`[1011001, 1011002]`, nil).Once()
	s.MarvelsUsecase.On("Search", mock.Anything, map[string]string{}).Return(initiateCharacterCollectionResponse(2, 2, 0), nil).Once()
	arr, merr := s.usecase.Search(context.Background())
	s.Assert().Equal(arr, []uint{1011001, 1011002})
	s.Assert().Nil(merr)
}

func (s *SearchTestSuite) TestCacheHitMarshalError() {
	s.CacheRepository.On("Get", mock.Anything, "marvels-characters").Return("zzz", nil).Once()
	s.MarvelsUsecase.On("Search", mock.Anything, map[string]string{}).Return(initiateCharacterCollectionResponse(2, 2, 0), nil).Once()
	_, merr := s.usecase.Search(context.Background())
	s.Assert().NotNil(merr)
}

func (s *SearchTestSuite) TestCallMarvelsError() {
	s.CacheRepository.On("Get", mock.Anything, "marvels-characters").Return("", nil).Once()
	s.MarvelsUsecase.On("Search", mock.Anything, map[string]string{"limit": "100", "offset": "0"}).Return(nil, errors.New("SomeError")).Once()
	_, merr := s.usecase.Search(context.Background())
	s.Assert().NotNil(merr)
}

func (s *SearchTestSuite) TestCallMarvelsSuccess() {
	s.CacheRepository.On("Get", mock.Anything, "marvels-characters").Return("", nil).Once()
	s.MarvelsUsecase.On("Search", mock.Anything, map[string]string{"limit": "100", "offset": "0"}).Return(initiateCharacterCollectionResponse(2, 2, 0), nil).Once()
	s.CacheRepository.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	arr, merr := s.usecase.Search(context.Background())
	s.Assert().Equal(arr, []uint{1011001, 1011002})
	s.Assert().Nil(merr)
}

func (s *SearchTestSuite) TestCallMarvelsSuccessMany() {
	s.CacheRepository.On("Get", mock.Anything, "marvels-characters").Return("", nil).Once()
	s.MarvelsUsecase.On("Search", mock.Anything, map[string]string{"limit": "100", "offset": "0"}).Return(initiateCharacterCollectionResponse(293, 100, 0), nil).Once()
	s.MarvelsUsecase.On("Search", mock.Anything, map[string]string{"limit": "100", "offset": "100"}).Return(initiateCharacterCollectionResponse(293, 100, 100), nil).Once()
	s.MarvelsUsecase.On("Search", mock.Anything, map[string]string{"limit": "100", "offset": "200"}).Return(initiateCharacterCollectionResponse(293, 93, 200), nil).Once()
	s.CacheRepository.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	arr, merr := s.usecase.Search(context.Background())
	s.Assert().Equal(len(arr), 293)
	s.Assert().Nil(merr)
}
