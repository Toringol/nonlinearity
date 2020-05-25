package story

import "github.com/Toringol/nonlinearity/app/model"

// Usecase - funcs interact with Story
type Usecase interface {
	SelectStoryByID(string) (*model.Story, error)
	SelectTopHeadings(map[string]string) (*model.StoryHeadings, error)
	UpdateStoryViews(string) error
	UpdateStoryRating(*model.RequestRating) (*model.Story, error)
	ListStories() ([]*model.Story, error)
	CreateStory(*model.Story) error
	UpdateStory(*model.Story) error
	DeleteStory(string) error
}
