package user

import "github.com/Toringol/nonlinearity/app/model"

// Usecase - funcs interact with User
type Usecase interface {
	// User usecase
	SelectUserByID(int64) (*model.User, error)
	SelectUserByUsername(string) (*model.User, error)
	CreateUser(*model.User) (int64, error)
	UpdateUser(*model.User) (int64, error)
	DeleteUser(int64) (int64, error)

	// Favourites usecase
	SelectUserFavouritesByID(int64) (*model.FavouriteCategories, error)
	CreateUserFavourites(int64, *model.FavouriteCategories) (int64, error)
	UpdateUserFavourites(int64, *model.FavouriteCategories) (int64, error)

	// Story usecase
	SelectStoryByID(int64) (*model.Story, error)
	SelectTopHeadingsStories(map[string]string) (*model.StoryHeadings, error)
	UpdateStoryViews(int64) (int64, error)
	UpdateStoryRating(*model.RequestRating) (int64, error)
	ListStories() ([]*model.Story, error)
	CreateStory(*model.Story) (int64, error)
	UpdateStory(*model.Story) (int64, error)
	DeleteStory(int64) (int64, error)

	// Genres usecase
	SelectGenresByStoryID(int64) ([]string, error)

	// RatingView usecase
	SelectView(int64, int64) (bool, error)
	SelectRate(int64, int64) (float64, bool, error)
	CreateView(*model.StoryRatingViews) (int64, error)
	UpdateRating(*model.StoryRatingViews) (int64, error)
}
