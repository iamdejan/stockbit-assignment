package repositories

import (
	"github.com/iamdejan/movie-search-service/internal/domain/pb"
	"net/http"
)

//go:generate mockgen -destination=../mocks/mock_repository.go -package=mocks github.com/iamdejan/movie-search-service/internal/repositories Repository
type Repository interface {
	Search(request *pb.MoviePreviewListRequest) (*pb.MoviePreviewListResponse, int, error)
	Get(id string) (*pb.Movie, int, error)
}

//go:generate mockgen -destination=../mocks/mock_http_client.go -package=mocks github.com/iamdejan/movie-search-service/internal/repositories HttpClient
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
