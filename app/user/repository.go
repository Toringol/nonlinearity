package user

import "github.com/Toringol/nonlinearity/app/model"

// Repository - funcs interact with DataBase
type Repository interface {
	// User table methods
	SelectUserByID(int64) (*model.User, error)
	SelectUserByUsername(string) (*model.User, error)
	CreateUser(*model.User) (int64, error)
	UpdateUser(*model.User) (int64, error)
	DeleteUser(int64) (int64, error)

	// Story table methods
	SelectStoryByID(int64) (*model.Story, error)
	SelectTopHeadingsStories(map[string]string) (*model.StoryHeadings, error)
	UpdateStoryViews(int64) (int64, error)
	UpdateStoryRating(*model.RequestRating) (int64, error)
	ListStories() ([]*model.Story, error)
	CreateStory(*model.Story) (int64, error)
	UpdateStory(*model.Story) (int64, error)
	DeleteStory(int64) (int64, error)
}
