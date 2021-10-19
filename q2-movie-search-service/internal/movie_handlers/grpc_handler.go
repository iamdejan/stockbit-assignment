package movie_handlers

import (
	"context"
	"github.com/iamdejan/movie-search-service/internal/domain/pb"
	"github.com/iamdejan/movie-search-service/internal/managers"
	"github.com/iamdejan/movie-search-service/internal/movie_manager"
)

type rpcHandler struct {
	manager managers.Manager
	pb.UnimplementedMovieSearchServiceServer
}

func NewRPCHandler() pb.MovieSearchServiceServer {
	return &rpcHandler{
		manager: movie_manager.NewMovieManager(),
	}
}

func (s *rpcHandler) Search(_ context.Context, r *pb.MoviePreviewListRequest) (*pb.MoviePreviewListResponse, error) {
	movieListResponse := s.manager.Search(r)
	if movieListResponse.Error() != nil {
		return nil, movieListResponse.Error()
	}

	return movieListResponse.Data(), nil
}

func (s *rpcHandler) Get(_ context.Context, r *pb.MovieRequest) (*pb.Movie, error) {
	movie := s.manager.Get(r.GetId())
	if movie.Error() != nil {
		return nil, movie.Error()
	}

	return movie.Data(), nil
}
