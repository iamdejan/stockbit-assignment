package main

import (
	"github.com/iamdejan/movie-search-service/config"
	"github.com/iamdejan/movie-search-service/internal/domain/pb"
	"github.com/iamdejan/movie-search-service/internal/movie_handlers"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"log"
	"net"
)

func startHTTPService() {
	e := echo.New()
	movie_handlers.RegisterRoutes(e, movie_handlers.NewRouteHandler())
	e.Logger.Fatal(e.Start(config.HttpPort))
}

func startRPCService() {
	listener, err := net.Listen("tcp", config.RpcPort)
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	pb.RegisterMovieSearchServiceServer(s, movie_handlers.NewRPCHandler())
	log.Printf("RPC server is listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	go startRPCService()
	startHTTPService()
}
