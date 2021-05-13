package handler_test

import (
	"errors"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/ivantedja/xmarvel/api/handler"
	"github.com/ivantedja/xmarvel/entity"
	"github.com/ivantedja/xmarvel/marvels/mocks"
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
	marvelsRepository := new(mocks.MarvelsRepository)
	handler := handler.NewMarvels(marvelsRepository, usecase)
	router := chi.NewRouter()
	router.Get("/characters", handler.Index)
	s.usecase = usecase
	s.handler = router
}

func (s *HandlerTestSuite) Record(request *http.Request) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	s.handler.ServeHTTP(response, request)
	return response
}

func (s *HandlerTestSuite) TestSuccessSearch() {
	s.usecase.On("Search", mock.Anything, map[string]string{"limit": "2", "modifiedSince": "2015-04-28"}).Return(&entity.CharacterCollection{}, nil).Once()
	r, _ := http.NewRequest("GET", "/characters", nil)
	w := s.Record(r)
	s.Assert().Equal(http.StatusOK, w.Code)
}

func (s *HandlerTestSuite) TestErrorSearch() {
	s.usecase.On("Search", mock.Anything, map[string]string{"limit": "2", "modifiedSince": "2015-04-28"}).Return(nil, errors.New("SomeError")).Once()
	r, _ := http.NewRequest("GET", "/characters", nil)
	w := s.Record(r)
	s.Assert().Equal(http.StatusBadRequest, w.Code)
}
