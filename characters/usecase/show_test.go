package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/ivantedja/xmarvel/characters"
)

type ShowTestSuite struct {
	suite.Suite
	usecase characters.Usecase
	BaseTestSuite
}

func TestShow(t *testing.T) {
	suite.Run(t, new(ShowTestSuite))
}

func (s *ShowTestSuite) SetupTest() {
	s.BaseTestSuite = getBaseTestSuite()
	s.usecase = initiateUsecase(&s.BaseTestSuite)
}

func (s *ShowTestSuite) TestCacheHitSuccess() {
	s.CacheRepository.On("Get", mock.Anything, "marvels-characters-1011000").Return(`{"id": 1011000,"name": "Abomination (Ultimate)","description": "Cool"}`, nil).Once()
	s.MarvelsUsecase.On("Show", mock.Anything, 1011000).Return(initiateCharacterResponse(), nil).Once()
	c, merr := s.usecase.Show(context.Background(), 1011000)
	s.Assert().Equal(c.ID, uint(1011000))
	s.Assert().Nil(merr)
}

func (s *ShowTestSuite) TestCacheHitMarshalError() {
	s.CacheRepository.On("Get", mock.Anything, "marvels-characters-1011000").Return("zzz", nil).Once()
	s.MarvelsUsecase.On("Show", mock.Anything, 1011000).Return(initiateCharacterResponse(), nil).Once()
	_, merr := s.usecase.Show(context.Background(), 1011000)
	s.Assert().NotNil(merr)
}

func (s *ShowTestSuite) TestCallMarvelsError() {
	s.CacheRepository.On("Get", mock.Anything, "marvels-characters-1011000").Return("", nil).Once()
	s.MarvelsUsecase.On("Show", mock.Anything, 1011000).Return(nil, errors.New("SomeError")).Once()
	_, merr := s.usecase.Show(context.Background(), 1011000)
	s.Assert().NotNil(merr)
}

func (s *ShowTestSuite) TestCallMarvelsSuccess() {
	s.CacheRepository.On("Get", mock.Anything, "marvels-characters-1011000").Return("", nil).Once()
	s.MarvelsUsecase.On("Show", mock.Anything, 1011000).Return(initiateCharacterResponse(), nil).Once()
	s.CacheRepository.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	c, merr := s.usecase.Show(context.Background(), 1011000)
	s.Assert().Equal(c.ID, uint(1011000))
	s.Assert().Nil(merr)
}
