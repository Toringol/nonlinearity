package http

import (
	"database/sql"
	"net/http"
	"strconv"

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

	// User handlers
	e.GET("/signin/", handlers.handleSignIn)
	e.GET("/profile/", handlers.handleGetUserProfile)
	e.GET("/logout/", handlers.handleLogout)

	e.POST("/signup/", handlers.handleSignUp)
	e.POST("/profile/", handlers.handleChangeUserProfile)
	e.POST("/changeAvatar/", handlers.handleChangeAvatar)

	// Story handlers
	e.GET("/getStory", handlers.handlerGetStoryInfo)
	e.GET("/topStories", handlers.handlerGetTopHeadings)

	e.POST("/endStory", handlers.handlerEndStory)
	e.POST("/rateStory", handlers.handlerRateStory)
}

// handleSignUp - create user record in DB if username is not occupied
// user`ll get default avatar from AWS S3 bucket
// setup session
func (h *userHandlers) handleSignUp(ctx echo.Context) error {
	userInput := new(model.User)

	if err := ctx.Bind(userInput); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	// Check user with input username in DB
	_, err := h.usecase.SelectUserByUsername(userInput.Username)
	switch {
	case err == sql.ErrNoRows:
		return echo.NewHTTPError(http.StatusConflict, "This username is occupied")
	case err != nil:
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	// Convert open pass to secure pass
	userInput.Password = string(tools.ConvertPass(userInput.Password))

	// Path to AWS S3 bucket and defaultAvatar
	userInput.Avatar = viper.GetString("storagePath") + "avatars/defaultAvatar"

	lastID, err := h.usecase.CreateUser(userInput)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	// Favourites - for new user its empty structure
	favourites := &model.FavouriteCategories{}

	_, err = h.usecase.CreateUserFavourites(lastID, favourites)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	userInput.ID = lastID

	_, err = cookies.SetSession(ctx, userInput)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Cookie set error")
	}

	return ctx.JSON(http.StatusCreated, userInput)
}

// handleSignIn - check user input
// if all ok -> setup session
func (h *userHandlers) handleSignIn(ctx echo.Context) error {
	authCredentials := new(model.User)

	if err := ctx.Bind(authCredentials); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	userRecord, err := h.usecase.SelectUserByUsername(authCredentials.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	if !tools.CheckPass(tools.ConvertPass(authCredentials.Password), userRecord.Password) {
		return echo.NewHTTPError(http.StatusUnauthorized, "Incorrect username or password")
	}

	_, err = cookies.SetSession(ctx, userRecord)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Cookie set error")
	}

	return ctx.JSON(http.StatusOK, authCredentials)
}

// handleGetUserProfile - check session and give data to user
func (h *userHandlers) handleGetUserProfile(ctx echo.Context) error {
	session, err := cookies.СheckSession(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	userData, err := h.usecase.SelectUserByUsername(session.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	userData.Favourited, err = h.usecase.SelectUserFavouritesByID(userData.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
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
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	oldUserData, err := h.usecase.SelectUserByUsername(session.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	changeUserData := new(model.User)

	if err := ctx.Bind(changeUserData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	changeUserData.ID = oldUserData.ID
	changeUserData.Avatar = oldUserData.Avatar

	if changeUserData.Username == "" {
		changeUserData.Username = oldUserData.Username
	} else if changeUserData.Password == "" {
		changeUserData.Password = oldUserData.Password
	}

	_, err = h.usecase.UpdateUser(changeUserData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	_, err = cookies.SetSession(ctx, changeUserData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Cookie set error")
	}

	return ctx.JSON(http.StatusCreated, "")
}

// handleChangeAvatar - check session if ok -> loadAvatar to AWS S3 bucket
// then change user`s avatar column in DB
func (h *userHandlers) handleChangeAvatar(ctx echo.Context) error {
	session, err := cookies.СheckSession(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	fileName, err := tools.LoadAvatar(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	oldUserData, err := h.usecase.SelectUserByUsername(session.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	oldUserData.Avatar = viper.GetString("storagePath") + fileName

	_, err = h.usecase.UpdateUser(oldUserData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	return ctx.JSON(http.StatusCreated, viper.GetString("storagePath")+fileName)
}

// handleLogout - delete session
func (h *userHandlers) handleLogout(ctx echo.Context) error {
	err := cookies.ClearSession(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Cookie del error")
	}

	return ctx.JSON(http.StatusOK, "")
}

// handlerGetStoryInfo - get story information from story DB
func (h *userHandlers) handlerGetStoryInfo(ctx echo.Context) error {

	idStr := ctx.QueryParam("id")
	if idStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "")
	}

	story, err := h.usecase.SelectStoryByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	// Get genres of story
	story.Genres, err = h.usecase.SelectGenresByStoryID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	return ctx.JSON(http.StatusOK, story)
}

// handlerGetTopHeadings - return top 10 stories in their headings
func (h *userHandlers) handlerGetTopHeadings(ctx echo.Context) error {

	headings := map[string]string{
		"Now popular":   "views",
		"Short stories": "description",
		"New":           "publicationDate",
	}

	storyHeadings, err := h.usecase.SelectTopHeadingsStories(headings)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	return ctx.JSON(http.StatusOK, storyHeadings)
}

// handlerEndStory - inc story views
func (h *userHandlers) handlerEndStory(ctx echo.Context) error {
	_, err := cookies.СheckSession(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	requestID := new(model.RequestIDStory)

	if err := ctx.Bind(requestID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	userID, err := strconv.ParseInt(cookies.ReadUserID(ctx).ID, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Error")
	}

	storyRatingViews := &model.StoryRatingViews{
		StoryID: requestID.ID,
		UserID:  userID,
		View:    true,
	}

	if _, err := h.usecase.CreateView(storyRatingViews); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	if _, err := h.usecase.UpdateStoryViews(requestID.ID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	genres, err := h.usecase.SelectGenresByStoryID(requestID.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	genresViewed := &model.FavouriteCategories{}

	for _, genre := range genres {
		switch genre {
		case "drama":
			genresViewed.Drama++
		case "romance":
			genresViewed.Romance++
		case "comedy":
			genresViewed.Comedy++
		case "horror":
			genresViewed.Horror++
		case "detective":
			genresViewed.Detective++
		case "fantasy":
			genresViewed.Fantasy++
		case "action":
			genresViewed.Action++
		case "realism":
			genresViewed.Realism++
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal Error")
		}
	}

	if _, err := h.usecase.UpdateUserFavourites(userID, genresViewed); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	return ctx.JSON(http.StatusOK, nil)
}

// handlerRateStory - update story rating and return new data of story
func (h *userHandlers) handlerRateStory(ctx echo.Context) error {
	_, err := cookies.СheckSession(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	reqRating := new(model.RequestRating)

	if err := ctx.Bind(reqRating); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	userID, err := strconv.ParseInt(cookies.ReadUserID(ctx).ID, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Error")
	}

	storyRatingViews := &model.StoryRatingViews{
		StoryID:      reqRating.ID,
		UserID:       userID,
		Rating:       true,
		PreviousRate: reqRating.Rating,
	}

	if _, err := h.usecase.UpdateRating(storyRatingViews); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	story, err := h.usecase.UpdateStoryRating(reqRating)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal DB Error")
	}

	return ctx.JSON(http.StatusOK, story)
}
