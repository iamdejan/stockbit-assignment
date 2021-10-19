package movie_response

import (
	"github.com/iamdejan/movie-search-service/internal/domain/pb"
	"net/http"
)

type MovieListResponse struct {
	statusCode int
	data       *pb.MoviePreviewListResponse
	err        error
}

func NewMovieListFailedResponse(statusCode int, err error) *MovieListResponse {
	return &MovieListResponse{
		statusCode: statusCode,
		data:       nil,
		err:        err,
	}
}

func NewMovieListSuccessResponse(data *pb.MoviePreviewListResponse) *MovieListResponse {
	return &MovieListResponse{
		statusCode: http.StatusOK,
		data:       data,
		err:        nil,
	}
}

func (m *MovieListResponse) StatusCode() int {
	return m.statusCode
}

func (m *MovieListResponse) Data() *pb.MoviePreviewListResponse {
	return m.data
}

func (m *MovieListResponse) Error() error {
	return m.err
}
