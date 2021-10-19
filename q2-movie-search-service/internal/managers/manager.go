package managers

import (
	"github.com/iamdejan/movie-search-service/internal/domain/movie_response"
	"github.com/iamdejan/movie-search-service/internal/domain/pb"
)

//go:generate mockgen -destination=../mocks/mock_manager.go -package=mocks github.com/iamdejan/movie-search-service/internal/managers Manager
type Manager interface {
	Search(request *pb.MoviePreviewListRequest) *movie_response.MovieListResponse
	Get(id string) *movie_response.MovieResponse
}
