package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/ivantedja/xmarvel/marvels"
)

type SearchTestSuite struct {
	suite.Suite
	usecase marvels.Usecase
	BaseTestSuite
}

func TestSearch(t *testing.T) {
	suite.Run(t, new(SearchTestSuite))
}

func (s *SearchTestSuite) SetupTest() {
	s.BaseTestSuite = getBaseTestSuite()
	s.usecase = initiateUsecase(&s.BaseTestSuite)
}

func (s *SearchTestSuite) TestSuccessSearch() {
	s.MarvelsRepository.On("Search", mock.Anything, map[string]string{}).Return(initiateCharacterCollectionResponse(), nil).Once()
	_, err := s.usecase.Search(context.Background(), map[string]string{})
	s.Assert().Nil(err)
}

func (s *SearchTestSuite) TestFailedSearch() {
	s.MarvelsRepository.On("Search", mock.Anything, map[string]string{}).Return(nil, errors.New("SomeError")).Once()
	_, err := s.usecase.Search(context.Background(), map[string]string{})
	s.Assert().NotNil(err)
}
