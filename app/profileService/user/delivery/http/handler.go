package http

import (
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
	return nil
}

func (h *userHandlers) handleSignIn(ctx echo.Context) error {
	return nil
}

func (h *userHandlers) handleGetUserProfile(ctx echo.Context) error {
	return nil
}

func (h *userHandlers) handleChangeUserProfile(ctx echo.Context) error {
	return nil
}
