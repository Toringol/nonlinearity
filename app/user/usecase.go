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

	// Story usecase
	SelectStoryByID(int64) (*model.Story, error)
	SelectTopHeadingsStories(map[string]string) (*model.StoryHeadings, error)
	UpdateStoryViews(int64) (int64, error)
	UpdateStoryRating(*model.RequestRating) (int64, error)
	ListStories() ([]*model.Story, error)
	CreateStory(*model.Story) (int64, error)
	UpdateStory(*model.Story) (int64, error)
	DeleteStory(int64) (int64, error)
}
