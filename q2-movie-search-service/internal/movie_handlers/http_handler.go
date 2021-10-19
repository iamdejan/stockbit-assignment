package movie_handlers

import (
	"github.com/iamdejan/movie-search-service/internal/domain/generic_error"
	"github.com/iamdejan/movie-search-service/internal/domain/pb"
	"github.com/iamdejan/movie-search-service/internal/handler"
	"github.com/iamdejan/movie-search-service/internal/managers"
	"github.com/iamdejan/movie-search-service/internal/movie_manager"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func RegisterRoutes(e *echo.Echo, h handler.Handler) {
	e.GET("/movies", h.Search)
	e.GET("/movies/:id", h.Get)
}

type routeHandler struct {
	manager managers.Manager
}

func NewRouteHandler() handler.Handler {
	return &routeHandler{
		manager: movie_manager.NewMovieManager(),
	}
}

func (h *routeHandler) Search(c echo.Context) error {
	paginationStr := c.QueryParam("pagination")
	pagination, err := strconv.ParseInt(paginationStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, generic_error.NewGenericError(err))
	}

	keyword := c.QueryParam("searchword")
	previewListRequest := &pb.MoviePreviewListRequest{
		Pagination: pagination,
		Searchword: keyword,
	}
	response := h.manager.Search(previewListRequest)
	if response.Error() != nil {
		return c.JSON(response.StatusCode(), generic_error.NewGenericError(response.Error()))
	}

	return c.JSON(response.StatusCode(), response.Data())
}

func (h *routeHandler) Get(c echo.Context) error {
	id := c.Param("id")
	response := h.manager.Get(id)
	if response.Error() != nil {
		return c.JSON(response.StatusCode(), generic_error.NewGenericError(response.Error()))
	}

	return c.JSON(response.StatusCode(), response.Data())
}
