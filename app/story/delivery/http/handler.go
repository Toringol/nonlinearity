package http

import (
	"net/http"

	"github.com/Toringol/nonlinearity/app/model"
	"github.com/Toringol/nonlinearity/app/story"
	"github.com/labstack/echo"
)

// storyHandlers - http handlers structure
type storyHandlers struct {
	usecase story.Usecase
}

// NewStoryHandler - deliver http handlers and listeners
func NewStoryHandler(e *echo.Echo, us story.Usecase) {
	handlers := storyHandlers{usecase: us}

	e.GET("/getStory", handlers.handlerGetStoryInfo)
	e.GET("/topStories", handlers.handlerGetTopHeadings)

	e.POST("/endStory", handlers.handlerEndStory)
	e.POST("/rateStory", handlers.handlerRateStory)
}

// handlerGetStoryInfo - get story information from story DB
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

// handlerGetTopHeadings - return top 10 stories in their headings
func (h *storyHandlers) handlerGetTopHeadings(ctx echo.Context) error {

	headings := map[string]string{
		"Now popular":   "views",
		"Short stories": "description",
		"New":           "publicationDate",
	}

	storyHeadings, err := h.usecase.SelectTopHeadings(headings)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	return ctx.JSON(http.StatusOK, storyHeadings)
}

// handlerEndStory - inc story views
func (h *storyHandlers) handlerEndStory(ctx echo.Context) error {

	requestID := new(model.RequestIDStory)

	if err := ctx.Bind(requestID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	if err := h.usecase.UpdateStoryViews(string(requestID.ID)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	return ctx.JSON(http.StatusOK, nil)
}

// handlerRateStory - update story rating and return new data of story
func (h *storyHandlers) handlerRateStory(ctx echo.Context) error {

	reqRating := new(model.RequestRating)

	if err := ctx.Bind(reqRating); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	story, err := h.usecase.UpdateStoryRating(reqRating)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	return ctx.JSON(http.StatusOK, story)
}
