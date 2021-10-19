package omdb_repo

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/iamdejan/movie-search-service/config"
	"github.com/iamdejan/movie-search-service/internal/domain/pb"
	"github.com/iamdejan/movie-search-service/internal/repositories"
	"net/http"
	"strconv"
	"time"
)

type omdbRepository struct {
	baseURL string
	client  repositories.HttpClient
}

func NewOMDBRepository() repositories.Repository {
	return &omdbRepository{
		baseURL: config.BaseURL,
		client:  &http.Client{Timeout: 3 * time.Second},
	}
}

func getBaseRequest(url string) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	q := request.URL.Query()
	q.Add("apikey", config.ApiKey)
	request.URL.RawQuery = q.Encode()
	return request, nil
}

func (o *omdbRepository) Search(r *pb.MoviePreviewListRequest) (*pb.MoviePreviewListResponse, int, error) {
	url := o.baseURL
	request, err := getBaseRequest(url)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	q := request.URL.Query()
	q.Add("s", r.GetSearchword())
	q.Add("page", strconv.FormatInt(r.GetPagination(), 10))
	request.URL.RawQuery = q.Encode()

	response, err := o.client.Do(request)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, response.StatusCode, fmt.Errorf(response.Status)
	}

	defer response.Body.Close()

	var responseBody = &pb.MoviePreviewListResponse{}
	if err := jsonpb.Unmarshal(response.Body, responseBody); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if responseBody.GetResponse() != "True" {
		return nil, http.StatusNotFound, fmt.Errorf(responseBody.GetError())
	}

	return responseBody, http.StatusOK, nil
}

func (o *omdbRepository) Get(id string) (*pb.Movie, int, error) {
	url := o.baseURL
	request, err := getBaseRequest(url)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	q := request.URL.Query()
	q.Add("i", id)
	request.URL.RawQuery = q.Encode()

	response, err := o.client.Do(request)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, response.StatusCode, fmt.Errorf(response.Status)
	}

	defer response.Body.Close()

	var responseBody = &pb.Movie{}
	if err := jsonpb.Unmarshal(response.Body, responseBody); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if responseBody.GetResponse() != "True" {
		return nil, http.StatusNotFound, fmt.Errorf(responseBody.GetError())
	}

	return responseBody, http.StatusOK, nil
}
