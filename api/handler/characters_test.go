package handler_test

import (
	"errors"
	"github.com/ivantedja/xmarvel/entity"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/ivantedja/xmarvel/api/handler"
	"github.com/ivantedja/xmarvel/characters/mocks"
)

type HandlerTestSuite struct {
	suite.Suite
	usecase *mocks.Usecase
	handler http.Handler
}

func TestHandler(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}

func (s *HandlerTestSuite) SetupTest() {
	usecase := new(mocks.Usecase)
	cacheRepository := new(mocks.CacheRepository)
	handler := handler.NewCharacters(cacheRepository, usecase)
	router := chi.NewRouter()
	router.Get("/characters", handler.Index)
	router.Get("/characters/{ID}", handler.Show)
	s.usecase = usecase
	s.handler = router
}

func (s *HandlerTestSuite) Record(request *http.Request) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	s.handler.ServeHTTP(response, request)
	return response
}

func (s *HandlerTestSuite) TestSuccessSearch() {
	s.usecase.On("Search", mock.Anything).Return([]uint{}, nil).Once()
	r, _ := http.NewRequest("GET", "/characters", nil)
	w := s.Record(r)
	s.Assert().Equal(http.StatusOK, w.Code)
}

func (s *HandlerTestSuite) TestErrorSearch() {
	s.usecase.On("Search", mock.Anything).Return([]uint{}, errors.New("SomeError")).Once()
	r, _ := http.NewRequest("GET", "/characters", nil)
	w := s.Record(r)
	s.Assert().Equal(http.StatusBadRequest, w.Code)
}

func (s *HandlerTestSuite) TestSuccessShow() {
	s.usecase.On("Show", mock.Anything, 1).Return(&entity.Character{}, nil).Once()
	r, _ := http.NewRequest("GET", "/characters/1", nil)
	w := s.Record(r)
	s.Assert().Equal(http.StatusOK, w.Code)
}

func (s *HandlerTestSuite) TestErrorParsingShow() {
	r, _ := http.NewRequest("GET", "/characters/a", nil)
	w := s.Record(r)
	s.Assert().Equal(http.StatusBadRequest, w.Code)
}

func (s *HandlerTestSuite) TestErrorShow() {
	s.usecase.On("Show", mock.Anything, 1).Return(&entity.Character{}, errors.New("SomeError")).Once()
	r, _ := http.NewRequest("GET", "/characters/1", nil)
	w := s.Record(r)
	s.Assert().Equal(http.StatusBadRequest, w.Code)
}
