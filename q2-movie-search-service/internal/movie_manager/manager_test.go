package movie_manager

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/iamdejan/movie-search-service/internal/domain/pb"
	"github.com/iamdejan/movie-search-service/internal/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewMovieManager(t *testing.T) {
	assert.NotNil(t, NewMovieManager())
}

func TestMovieManager_ListWithResult(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockRepository(mockCtrl)

	request := &pb.MoviePreviewListRequest{
		Pagination: 0,
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
	mockRepo.EXPECT().Search(gomock.Eq(request)).Return(moviePreviewResponse, http.StatusOK, nil)

	manager := &movieManager{
		repository: mockRepo,
	}
	movieListResponse := manager.Search(request)
	assert.NotNil(t, movieListResponse)
	assert.NotNil(t, movieListResponse.StatusCode())
	assert.Nil(t, movieListResponse.Error())
	assert.NotNil(t, movieListResponse.Data())
	assert.Equal(t, moviePreviewList, movieListResponse.Data().GetSearch())
}

func TestMovieManager_ListWithError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockRepository(mockCtrl)

	request := &pb.MoviePreviewListRequest{
		Pagination: 0,
		Searchword: "batman",
	}

	mockRepo.EXPECT().Search(gomock.Eq(request)).Return(nil, http.StatusInternalServerError, fmt.Errorf("internal server error"))

	manager := &movieManager{
		repository: mockRepo,
	}
	movieListResponse := manager.Search(request)
	assert.NotNil(t, movieListResponse)
	assert.NotNil(t, movieListResponse.StatusCode())
	assert.NotNil(t, movieListResponse.Error())
}

func TestMovieManager_GetWithResult(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockRepository(mockCtrl)

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
	mockRepo.EXPECT().Get(gomock.Eq(imdbId)).Return(movie, http.StatusOK, nil)

	manager := &movieManager{
		repository: mockRepo,
	}
	movieResponse := manager.Get(imdbId)
	assert.NotNil(t, movieResponse)
	assert.NotNil(t, movieResponse.StatusCode())
	assert.Nil(t, movieResponse.Error())
	assert.Equal(t, movie, movieResponse.Data())
}

func TestMovieManager_GetWithError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mocks.NewMockRepository(mockCtrl)

	imdbId := "tt123456"
	mockRepo.EXPECT().Get(gomock.Eq(imdbId)).Return(nil, http.StatusInternalServerError, fmt.Errorf("internal server error"))

	manager := &movieManager{
		repository: mockRepo,
	}
	movieResponse := manager.Get(imdbId)
	assert.NotNil(t, movieResponse)
	assert.NotNil(t, movieResponse.StatusCode())
	assert.NotNil(t, movieResponse.Error())
}
