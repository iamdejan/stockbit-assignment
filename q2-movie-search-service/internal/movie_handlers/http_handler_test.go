package movie_handlers

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/iamdejan/movie-search-service/internal/domain/generic_error"
	"github.com/iamdejan/movie-search-service/internal/domain/movie_response"
	"github.com/iamdejan/movie-search-service/internal/domain/pb"
	"github.com/iamdejan/movie-search-service/internal/mocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRegisterRoutes(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	e := echo.New()

	mockHandler := mocks.NewMockHandler(mockCtrl)
	RegisterRoutes(e, mockHandler)
}

func TestNewRouteHandler(t *testing.T) {
	assert.NotNil(t, NewRouteHandler())
}

func TestRouteHandler_SearchWithResult(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockContext := mocks.NewMockContext(mockCtrl)
	mockManager := mocks.NewMockManager(mockCtrl)

	h := &routeHandler{
		manager: mockManager,
	}

	moviePreview := &pb.MoviePreview{
		Title:  "Superman Returns",
		Year:   "2021",
		ImdbID: "tt123456",
		Type:   "Action",
		Poster: "",
	}
	moviePreviewList := make([]*pb.MoviePreview, 1)
	moviePreviewList[0] = moviePreview
	moviePreviewResponse := &pb.MoviePreviewListResponse{
		Search:       moviePreviewList,
		TotalResults: "1",
		Response:     "True",
		Error:        "",
	}
	response := movie_response.NewMovieListSuccessResponse(moviePreviewResponse)
	gomock.InOrder(
		mockContext.EXPECT().QueryParam(gomock.Eq("pagination")).Return("1"),
		mockContext.EXPECT().QueryParam(gomock.Eq("searchword")).Return("superman"),
		mockManager.EXPECT().Search(gomock.Any()).Return(response),
		mockContext.EXPECT().JSON(http.StatusOK, gomock.AssignableToTypeOf(response.Data())).Return(nil),
	)

	err := h.Search(mockContext)
	assert.Nil(t, err)
}

func TestRouteHandler_SearchWithQueryParamError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockContext := mocks.NewMockContext(mockCtrl)
	mockManager := mocks.NewMockManager(mockCtrl)

	h := &routeHandler{
		manager: mockManager,
	}

	gomock.InOrder(
		mockContext.EXPECT().QueryParam(gomock.Eq("pagination")).Return("x"),
		mockContext.EXPECT().JSON(gomock.Not(http.StatusOK), gomock.AssignableToTypeOf(generic_error.GenericError{})).Return(nil),
	)

	err := h.Search(mockContext)
	assert.Nil(t, err)
}

func TestRouteHandler_SearchWithManagerError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockContext := mocks.NewMockContext(mockCtrl)
	mockManager := mocks.NewMockManager(mockCtrl)

	h := &routeHandler{
		manager: mockManager,
	}

	response := movie_response.NewMovieListFailedResponse(http.StatusInternalServerError, fmt.Errorf("internal server error"))
	gomock.InOrder(
		mockContext.EXPECT().QueryParam(gomock.Eq("pagination")).Return("1"),
		mockContext.EXPECT().QueryParam(gomock.Eq("searchword")).Return("superman"),
		mockManager.EXPECT().Search(gomock.Any()).Return(response),
		mockContext.EXPECT().JSON(gomock.Eq(response.StatusCode()), gomock.AssignableToTypeOf(generic_error.GenericError{})).Return(nil),
	)

	err := h.Search(mockContext)
	assert.Nil(t, err)
}

func TestRouteHandler_GetWithResult(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockContext := mocks.NewMockContext(mockCtrl)
	mockManager := mocks.NewMockManager(mockCtrl)

	h := &routeHandler{
		manager: mockManager,
	}

	imdbId := "tt123456"
	movie := &pb.Movie{
		Title:    "Batman Returns",
		Year:     "2021",
		ImdbID:   imdbId,
		Type:     "Action",
		Poster:   "",
		Response: "True",
		Error:    "",
	}
	response := movie_response.NewMovieSuccessResponse(movie)
	gomock.InOrder(
		mockContext.EXPECT().Param(gomock.Eq("id")).Return(imdbId),
		mockManager.EXPECT().Get(gomock.Any()).Return(response),
		mockContext.EXPECT().JSON(http.StatusOK, gomock.AssignableToTypeOf(response.Data())).Return(nil),
	)

	err := h.Get(mockContext)
	assert.Nil(t, err)
}

func TestRouteHandler_GetWithError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockContext := mocks.NewMockContext(mockCtrl)
	mockManager := mocks.NewMockManager(mockCtrl)

	h := &routeHandler{
		manager: mockManager,
	}

	imdbId := "tt123456"
	response := movie_response.NewMovieFailedResponse(http.StatusInternalServerError, fmt.Errorf("internal server error"))
	gomock.InOrder(
		mockContext.EXPECT().Param(gomock.Eq("id")).Return(imdbId),
		mockManager.EXPECT().Get(gomock.Any()).Return(response),
		mockContext.EXPECT().JSON(gomock.Eq(response.StatusCode()), gomock.AssignableToTypeOf(generic_error.GenericError{})).Return(nil),
	)

	err := h.Get(mockContext)
	assert.Nil(t, err)
}
