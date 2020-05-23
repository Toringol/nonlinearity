package http

import (
	"net/http"

	"github.com/Toringol/nonlinearity/app/story"
	"github.com/labstack/echo"
)

type storyHandlers struct {
	usecase story.Usecase
}

func NewStoryHandler(e *echo.Echo, us story.Usecase) {
	handlers := storyHandlers{usecase: us}

	e.GET("/getStory", handlers.handlerGetStoryInfo)
}

func (h *storyHandlers) handlerGetStoryInfo(ctx echo.Context) error {

	id := ctx.QueryParam("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	story, err := h.usecase.SelectStoryByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	return ctx.JSON(http.StatusOK, story)
}
