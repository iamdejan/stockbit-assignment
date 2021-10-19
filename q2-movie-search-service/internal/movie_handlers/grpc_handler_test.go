package movie_handlers

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/iamdejan/movie-search-service/internal/domain/movie_response"
	"github.com/iamdejan/movie-search-service/internal/domain/pb"
	"github.com/iamdejan/movie-search-service/internal/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewRPCHandler(t *testing.T) {
	assert.NotNil(t, NewRPCHandler())
}

func TestRpcHandler_SearchWithResult(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockManager := mocks.NewMockManager(mockCtrl)

	s := &rpcHandler{
		manager: mockManager,
	}

	request := &pb.MoviePreviewListRequest{
		Pagination: 1,
		Searchword: "batman",
	}

	moviePreview := &pb.MoviePreview{
		Title:  "Batman Returns",
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
	mockResponse := movie_response.NewMovieListSuccessResponse(moviePreviewResponse)
	mockManager.EXPECT().Search(gomock.Eq(request)).Return(mockResponse)

	response, err := s.Search(context.TODO(), request)
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.Equal(t, moviePreviewList, response.GetSearch())
}

func TestRpcHandler_SearchWithError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockManager := mocks.NewMockManager(mockCtrl)

	s := &rpcHandler{
		manager: mockManager,
	}

	request := &pb.MoviePreviewListRequest{
		Pagination: 1,
		Searchword: "batman",
	}

	mockResponse := movie_response.NewMovieListFailedResponse(http.StatusInternalServerError, fmt.Errorf("internal server error"))
	mockManager.EXPECT().Search(gomock.Eq(request)).Return(mockResponse)

	response, err := s.Search(context.TODO(), request)
	assert.Nil(t, response)
	assert.Equal(t, mockResponse.Error(), err)
}

func TestRpcHandler_GetWithResult(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockManager := mocks.NewMockManager(mockCtrl)

	s := &rpcHandler{
		manager: mockManager,
	}

	imdbId := "tt123456"
	movieRequest := &pb.MovieRequest{Id: imdbId}
	movie := &pb.Movie{
		Title:    "Batman Returns",
		Year:     "2021",
		ImdbID:   imdbId,
		Type:     "Action",
		Poster:   "",
		Response: "True",
		Error:    "",
	}
	mockResponse := movie_response.NewMovieSuccessResponse(movie)

	mockManager.EXPECT().Get(gomock.Eq(movieRequest.GetId())).Return(mockResponse)

	response, err := s.Get(context.TODO(), movieRequest)
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.Equal(t, movie, response)
}

func TestRpcHandler_GetWithError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockManager := mocks.NewMockManager(mockCtrl)

	s := &rpcHandler{
		manager: mockManager,
	}

	imdbId := "tt123456"
	movieRequest := &pb.MovieRequest{Id: imdbId}
	mockResponse := movie_response.NewMovieFailedResponse(http.StatusInternalServerError, fmt.Errorf("internal server error"))

	mockManager.EXPECT().Get(gomock.Eq(movieRequest.GetId())).Return(mockResponse)

	response, err := s.Get(context.TODO(), movieRequest)
	assert.Nil(t, response)
	assert.Equal(t, mockResponse.Error(), err)
}
