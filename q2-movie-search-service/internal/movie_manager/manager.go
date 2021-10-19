package movie_manager

import (
	"github.com/iamdejan/movie-search-service/internal/domain/movie_response"
	"github.com/iamdejan/movie-search-service/internal/domain/pb"
	"github.com/iamdejan/movie-search-service/internal/managers"
	"github.com/iamdejan/movie-search-service/internal/omdb_repo"
	"github.com/iamdejan/movie-search-service/internal/repositories"
)

func NewMovieManager() managers.Manager {
	return &movieManager{
		repository: omdb_repo.NewOMDBRepository(),
	}
}

type movieManager struct {
	repository repositories.Repository
}

func (m *movieManager) Search(request *pb.MoviePreviewListRequest) *movie_response.MovieListResponse {
	moviePreviewResponse, statusCode, err := m.repository.Search(request)
	if err != nil {
		return movie_response.NewMovieListFailedResponse(statusCode, err)
	}

	return movie_response.NewMovieListSuccessResponse(moviePreviewResponse)
}

func (m *movieManager) Get(id string) *movie_response.MovieResponse {
	movie, statusCode, err := m.repository.Get(id)
	if err != nil {
		return movie_response.NewMovieFailedResponse(statusCode, err)
	}

	return movie_response.NewMovieSuccessResponse(movie)
}
