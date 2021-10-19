package movie_response

import (
	"github.com/iamdejan/movie-search-service/internal/domain/pb"
	"net/http"
)

type MovieResponse struct {
	statusCode int
	data       *pb.Movie
	err        error
}

func NewMovieFailedResponse(statusCode int, err error) *MovieResponse {
	return &MovieResponse{
		statusCode: statusCode,
		data:       nil,
		err:        err,
	}
}

func NewMovieSuccessResponse(data *pb.Movie) *MovieResponse {
	return &MovieResponse{
		statusCode: http.StatusOK,
		data:       data,
		err:        nil,
	}
}

func (m *MovieResponse) StatusCode() int {
	return m.statusCode
}

func (m *MovieResponse) Data() *pb.Movie {
	return m.data
}

func (m *MovieResponse) Error() error {
	return m.err
}
