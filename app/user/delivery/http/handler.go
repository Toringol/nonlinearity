package http

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/Toringol/nonlinearity/app/auth/cookies"
	"github.com/Toringol/nonlinearity/tools"

	"github.com/Toringol/nonlinearity/app/model"
	"github.com/Toringol/nonlinearity/app/user"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
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
	e.GET("/logout/", handlers.handleLogout)

	e.POST("/signup/", handlers.handleSignUp)
	e.POST("/profile/", handlers.handleChangeUserProfile)
	e.POST("/changeAvatar/", handlers.handleChangeAvatar)
}

// handleSignUp - create user record in DB if username is not occupied
// user`ll get default avatar from AWS S3 bucket
// setup session
func (h *userHandlers) handleSignUp(ctx echo.Context) error {
	userInput := new(model.User)

	if err := ctx.Bind(userInput); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request")
	}

	// Check user with input username in DB
	_, err := h.usecase.SelectUserByUsername(userInput.Username)
	switch {
	case err == sql.ErrNoRows:
		return ctx.JSON(http.StatusConflict, "User with this username already exist")
	case err != nil:
		return ctx.JSON(http.StatusInternalServerError, "Internal Error")
	}

	// Path to AWS S3 bucket and defaulAvatar
	userInput.Avatar = viper.GetString("imageStoragePath") + "avatars/defaultAvatar"

	lastID, err := h.usecase.CreateUser(userInput)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Error")
	}

	userInput.ID = lastID

	cookie, err := cookies.SetSession(ctx, userInput)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Cookie set error")
	}

	log.Println(cookie)

	return ctx.JSON(http.StatusOK, userInput)
}

// handleSignIn - check user input
// if all ok -> setup session
func (h *userHandlers) handleSignIn(ctx echo.Context) error {
	authCredentials := new(model.User)

	if err := ctx.Bind(authCredentials); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request")
	}

	userRecord, err := h.usecase.SelectUserByUsername(authCredentials.Username)
	if err != nil || authCredentials.Password != userRecord.Password {
		return ctx.JSON(http.StatusUnauthorized, "Incorrect username or password!")
	}

	cookie, err := cookies.SetSession(ctx, userRecord)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Error")
	}

	log.Println(cookie)

	return ctx.JSON(http.StatusOK, authCredentials)
}

// handleGetUserProfile - check session and give data to user
func (h *userHandlers) handleGetUserProfile(ctx echo.Context) error {
	session, err := cookies.СheckSession(ctx)
	if err != nil {
		return nil
	}

	userData, err := h.usecase.SelectUserByUsername(session.Username)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Error")
	}

	userData.ID = 0
	userData.Password = ""

	return ctx.JSON(http.StatusOK, userData)
}

// handleChangeUserProfile - check session then get old information from DB
// if some information user trying to change we replace it
// then update record in DB with new data
func (h *userHandlers) handleChangeUserProfile(ctx echo.Context) error {
	session, err := cookies.СheckSession(ctx)
	if err != nil {
		return nil
	}

	oldUserData, err := h.usecase.SelectUserByUsername(session.Username)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Error")
	}

	changeUserData := new(model.User)

	if err := ctx.Bind(changeUserData); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Bad Request")
	}

	changeUserData.ID = oldUserData.ID
	changeUserData.Avatar = oldUserData.Avatar
	nullTime := time.Time{}

	if changeUserData.Username == "" {
		changeUserData.Username = oldUserData.Username
	} else if changeUserData.Password == "" {
		changeUserData.Password = oldUserData.Password
	} else if changeUserData.UserPersonalData.DateOfBirth == nullTime {
		changeUserData.UserPersonalData.DateOfBirth = oldUserData.UserPersonalData.DateOfBirth
	} else if changeUserData.UserPersonalData.Relationship == "" {
		changeUserData.UserPersonalData.Relationship = oldUserData.UserPersonalData.Relationship
	} else if changeUserData.UserPersonalData.Status == "" {
		changeUserData.UserPersonalData.Status = oldUserData.UserPersonalData.Status
	}

	affected, err := h.usecase.UpdateUser(changeUserData)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Error")
	}

	log.Println("Update affectedRows: ", affected)

	cookie, err := cookies.SetSession(ctx, changeUserData)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Error")
	}

	log.Println(cookie)

	return nil
}

// handleChangeAvatar - check session if ok -> loadAvatar to AWS S3 bucket
// then change user`s avatar column in DB
func (h *userHandlers) handleChangeAvatar(ctx echo.Context) error {
	session, err := cookies.СheckSession(ctx)
	if err != nil {
		return nil
	}

	fileName, err := tools.LoadAvatar(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Error")
	}

	oldUserData, err := h.usecase.SelectUserByUsername(session.Username)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Error")
	}

	oldUserData.Avatar = viper.GetString("imageStoragePath") + fileName

	_, err = h.usecase.UpdateUser(oldUserData)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Error")
	}

	return ctx.JSON(http.StatusOK, viper.GetString("imageStoragePath")+fileName)
}

// handleLogout - delete session
func (h *userHandlers) handleLogout(ctx echo.Context) error {
	err := cookies.ClearSession(ctx)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Internal Error")
	}

	return ctx.JSON(http.StatusOK, "")
}
