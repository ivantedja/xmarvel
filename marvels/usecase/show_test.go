package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/ivantedja/xmarvel/marvels"
)

type ShowTestSuite struct {
	suite.Suite
	usecase marvels.Usecase
	BaseTestSuite
}

func TestShow(t *testing.T) {
	suite.Run(t, new(ShowTestSuite))
}

func (s *ShowTestSuite) SetupTest() {
	s.BaseTestSuite = getBaseTestSuite()
	s.usecase = initiateUsecase(&s.BaseTestSuite)
}

func (s *ShowTestSuite) TestSuccessShow() {
	s.MarvelsRepository.On("Show", mock.Anything, 1016823).Return(initiateCharacterResponse(), nil).Once()
	_, err := s.usecase.Show(context.Background(), 1016823)
	s.Assert().Nil(err)
}

func (s *ShowTestSuite) TestFailedShow() {
	s.MarvelsRepository.On("Show", mock.Anything, 9999999999).Return(nil, errors.New("SomeError")).Once()
	_, err := s.usecase.Show(context.Background(), 9999999999)
	s.Assert().NotNil(err)
}
