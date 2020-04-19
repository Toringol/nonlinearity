package http

import (
	"net/http"

	"github.com/Toringol/nonlinearity/app/profileService/model"
	"github.com/Toringol/nonlinearity/app/profileService/user"
	"github.com/labstack/echo"
)

// userHandlers - http handlers structure
type userHandlers struct {
	usecase user.Usecase
}

// NewUserHandler - deliver our handlers in http
func NewUserHandler(e *echo.Echo, us user.Usecase) {
	handlers := userHandlers{usecase: us}

	e.GET("/signin/", handlers.handleSignIn)
	e.GET("/profile/", handlers.handleGetUserProfile)

	e.POST("/signup/", handlers.handleSignUp)
	e.POST("/profile/", handlers.handleChangeUserProfile)
}

func (h *userHandlers) handleSignUp(ctx echo.Context) error {
	userInput := new(model.User)

	if err := ctx.Bind(userInput); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request")
	}

	userInput.Avatar = "default" // TODO: change it by adding normal path

	lastID, err := h.usecase.CreateUser(userInput)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Error")
	}

	_ = lastID // TODO: use lastID for SesstionToken

	return ctx.JSON(http.StatusOK, userInput)
}

func (h *userHandlers) handleSignIn(ctx echo.Context) error {
	authCredentials := new(model.User)

	if err := ctx.Bind(authCredentials); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request")
	}

	userRecord, err := h.usecase.SelectUserbyUsername(authCredentials.Username)
	if err != nil || authCredentials.Password != userRecord.Password {
		return ctx.JSON(http.StatusUnauthorized, "Incorrect username or password!")
	}

	return ctx.JSON(http.StatusOK, authCredentials)
}

func (h *userHandlers) handleGetUserProfile(ctx echo.Context) error {

	return nil
}

func (h *userHandlers) handleChangeUserProfile(ctx echo.Context) error {
	return nil
}
