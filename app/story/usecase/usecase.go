package usecase

import (
	"github.com/Toringol/nonlinearity/app/model"
	"github.com/Toringol/nonlinearity/app/story"
)

// storyUsecase - connector database and handlers
type storyUsecase struct {
	repo story.Repository
}

// NewStoryUsecase - create new connector between handlers and story DB
func NewStoryUsecase(storyRepo story.Repository) story.Usecase {
	return storyUsecase{repo: storyRepo}
}

// SelectStoryByID - return story`s data by ID
func (su storyUsecase) SelectStoryByID(id string) (*model.Story, error) {
	return su.repo.SelectByID(id)
}

// SelectTopHeadings - return top 10 stories for each heading
func (su storyUsecase) SelectTopHeadings(headings map[string]string) (*model.StoryHeadings, error) {
	return su.repo.SelectTopHeadings(headings)
}

// UpdateStoryViews - inc story views
func (su storyUsecase) UpdateStoryViews(id string) error {
	return su.repo.UpdateViews(id)
}

// UpdateStoryRating - update story rating
func (su storyUsecase) UpdateStoryRating(reqRating *model.RequestRating) (*model.Story, error) {
	return su.repo.UpdateRating(reqRating)
}

// ListStories - return all info about all stories in story DB
func (su storyUsecase) ListStories() ([]*model.Story, error) {
	return su.repo.List()
}

// CreateStory - create new story record in story DB
func (su storyUsecase) CreateStory(story *model.Story) error {
	return su.repo.Create(story)
}

// UpdateStory - update records in story DB
func (su storyUsecase) UpdateStory(story *model.Story) error {
	return su.repo.Update(story)
}

// DeleteStory - delete story record from story DB
func (su storyUsecase) DeleteStory(id string) error {
	return su.repo.Delete(id)
}
