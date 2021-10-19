package omdb_repo

import (
	"bytes"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/jsonpb"
	"github.com/iamdejan/movie-search-service/config"
	"github.com/iamdejan/movie-search-service/internal/domain/pb"
	"github.com/iamdejan/movie-search-service/internal/mocks"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestNewOMDBRepository(t *testing.T) {
	assert.NotNil(t, NewOMDBRepository())
}

func TestOmdbRepository_Search(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := mocks.NewMockHttpClient(mockCtrl)

	repository := &omdbRepository{
		baseURL: config.BaseURL,
		client:  mockClient,
	}

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

	marshaler := jsonpb.Marshaler{}
	jsonStr, err := marshaler.MarshalToString(moviePreviewResponse)
	if err != nil {
		t.Fail()
	}

	httpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(jsonStr))),
	}
	mockClient.EXPECT().Do(gomock.AssignableToTypeOf(&http.Request{})).Return(httpResponse, nil)

	actualResponse, statusCode, err := repository.Search(request)
	assert.Nil(t, err)
	assert.Equal(t, httpResponse.StatusCode, statusCode)
	assert.Equal(t, moviePreviewResponse, actualResponse)
}

func TestOmdbRepository_SearchWithError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := mocks.NewMockHttpClient(mockCtrl)

	repository := &omdbRepository{
		baseURL: config.BaseURL,
		client:  mockClient,
	}

	request := &pb.MoviePreviewListRequest{
		Pagination: 0,
		Searchword: "batman",
	}

	mockClient.EXPECT().Do(gomock.AssignableToTypeOf(&http.Request{})).Return(nil, fmt.Errorf("internal server error"))

	actualResponse, statusCode, err := repository.Search(request)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, statusCode)
	assert.Nil(t, actualResponse)
}

func TestOmdbRepository_SearchButUnauthorized(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := mocks.NewMockHttpClient(mockCtrl)

	repository := &omdbRepository{
		baseURL: config.BaseURL,
		client:  mockClient,
	}

	request := &pb.MoviePreviewListRequest{
		Pagination: 0,
		Searchword: "batman",
	}

	httpResponse := &http.Response{
		StatusCode: http.StatusUnauthorized,
	}
	mockClient.EXPECT().Do(gomock.AssignableToTypeOf(&http.Request{})).Return(httpResponse, nil)

	actualResponse, statusCode, err := repository.Search(request)
	assert.NotNil(t, err)
	assert.Equal(t, httpResponse.StatusCode, statusCode)
	assert.Nil(t, actualResponse)
}

func TestOmdbRepository_SearchButFailToParse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := mocks.NewMockHttpClient(mockCtrl)

	repository := &omdbRepository{
		baseURL: config.BaseURL,
		client:  mockClient,
	}

	request := &pb.MoviePreviewListRequest{
		Pagination: 0,
		Searchword: "batman",
	}

	jsonStr := `{"name":"James","full_name":"James Bond"}`
	httpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(jsonStr))),
	}
	mockClient.EXPECT().Do(gomock.AssignableToTypeOf(&http.Request{})).Return(httpResponse, nil)

	actualResponse, statusCode, err := repository.Search(request)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, statusCode)
	assert.Nil(t, actualResponse)
}

func TestOmdbRepository_SearchButNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := mocks.NewMockHttpClient(mockCtrl)

	repository := &omdbRepository{
		baseURL: config.BaseURL,
		client:  mockClient,
	}

	request := &pb.MoviePreviewListRequest{
		Pagination: 0,
		Searchword: "batman",
	}

	moviePreviewResponse := &pb.MoviePreviewListResponse{
		Response: "False",
		Error:    "Incorrect IMDb ID.",
	}

	marshaler := jsonpb.Marshaler{}
	jsonStr, err := marshaler.MarshalToString(moviePreviewResponse)
	if err != nil {
		t.Fail()
	}

	httpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(jsonStr))),
	}
	mockClient.EXPECT().Do(gomock.AssignableToTypeOf(&http.Request{})).Return(httpResponse, nil)

	actualResponse, statusCode, err := repository.Search(request)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, statusCode)
	assert.Nil(t, actualResponse)
}

func TestOmdbRepository_Get(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := mocks.NewMockHttpClient(mockCtrl)

	repository := &omdbRepository{
		baseURL: config.BaseURL,
		client:  mockClient,
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

	marshaler := jsonpb.Marshaler{}
	jsonStr, err := marshaler.MarshalToString(movie)
	if err != nil {
		t.Fail()
	}

	httpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(jsonStr))),
	}
	mockClient.EXPECT().Do(gomock.AssignableToTypeOf(&http.Request{})).Return(httpResponse, nil)

	actualResponse, statusCode, err := repository.Get(imdbId)
	assert.Nil(t, err)
	assert.Equal(t, httpResponse.StatusCode, statusCode)
	assert.Equal(t, movie, actualResponse)
}

func TestOmdbRepository_GetWithError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := mocks.NewMockHttpClient(mockCtrl)

	repository := &omdbRepository{
		baseURL: config.BaseURL,
		client:  mockClient,
	}

	imdbId := "tt123456"
	mockClient.EXPECT().Do(gomock.AssignableToTypeOf(&http.Request{})).Return(nil, fmt.Errorf("internal server error"))

	actualResponse, statusCode, err := repository.Get(imdbId)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, statusCode)
	assert.Nil(t, actualResponse)
}

func TestOmdbRepository_GetButUnauthorized(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := mocks.NewMockHttpClient(mockCtrl)

	repository := &omdbRepository{
		baseURL: config.BaseURL,
		client:  mockClient,
	}

	imdbId := "tt123456"

	httpResponse := &http.Response{
		StatusCode: http.StatusUnauthorized,
		Body:       nil,
	}
	mockClient.EXPECT().Do(gomock.AssignableToTypeOf(&http.Request{})).Return(httpResponse, nil)

	actualResponse, statusCode, err := repository.Get(imdbId)
	assert.NotNil(t, err)
	assert.Equal(t, httpResponse.StatusCode, statusCode)
	assert.Nil(t, actualResponse)
}

func TestOmdbRepository_GetButFailToParse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := mocks.NewMockHttpClient(mockCtrl)

	repository := &omdbRepository{
		baseURL: config.BaseURL,
		client:  mockClient,
	}

	imdbId := "tt123456"
	jsonStr := `{"name":"Name","full_name":"test f name"}`

	httpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(jsonStr))),
	}
	mockClient.EXPECT().Do(gomock.AssignableToTypeOf(&http.Request{})).Return(httpResponse, nil)

	actualResponse, statusCode, err := repository.Get(imdbId)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusInternalServerError, statusCode)
	assert.Nil(t, actualResponse)
}

func TestOmdbRepository_GetButNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockClient := mocks.NewMockHttpClient(mockCtrl)

	repository := &omdbRepository{
		baseURL: config.BaseURL,
		client:  mockClient,
	}

	imdbId := "tt123456"
	movie := &pb.Movie{
		Response: "False",
		Error:    "Incorrect IMDb ID.",
	}

	marshaler := jsonpb.Marshaler{}
	jsonStr, err := marshaler.MarshalToString(movie)
	if err != nil {
		t.Fail()
	}

	httpResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(jsonStr))),
	}
	mockClient.EXPECT().Do(gomock.AssignableToTypeOf(&http.Request{})).Return(httpResponse, nil)

	actualResponse, statusCode, err := repository.Get(imdbId)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, statusCode)
	assert.Nil(t, actualResponse)
}
