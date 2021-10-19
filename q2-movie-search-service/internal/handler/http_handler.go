package handler

import (
	"github.com/labstack/echo/v4"
)

//go:generate mockgen -destination=../mocks/mock_handler.go -package=mocks github.com/iamdejan/movie-search-service/internal/handler Handler
type Handler interface {
	Search(c echo.Context) error
	Get(c echo.Context) error
}
