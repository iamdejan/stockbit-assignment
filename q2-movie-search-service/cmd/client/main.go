package main

import (
	"context"
	"fmt"
	"github.com/iamdejan/movie-search-service/config"
	"github.com/iamdejan/movie-search-service/internal/domain/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

func showMoviePreviewList(ctx context.Context, c pb.MovieSearchServiceClient, keyword string) {
	r, err := c.Search(ctx, &pb.MoviePreviewListRequest{
		Pagination: 1,
		Searchword: keyword,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("list of movies with title '%s': {%v}\n", keyword, r)
}

func showMovie(ctx context.Context, c pb.MovieSearchServiceClient, id string) {
	r, err := c.Get(ctx, &pb.MovieRequest{
		Id: id,
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("movie with id '%s': {%v}\n", id, r)
}

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("localhost%s", config.RpcPort), grpc.WithInsecure(), grpc.FailOnNonTempDialError(true), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	c := pb.NewMovieSearchServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	showMoviePreviewList(ctx, c, "superman")
	showMovie(ctx, c, "tt0372784")
}
