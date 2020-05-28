package usecase

import (
	"github.com/Toringol/nonlinearity/app/model"
	"github.com/Toringol/nonlinearity/app/user"
)

// userUsecase - connector database and user
type userUsecase struct {
	repo user.Repository
}

// NewUserUsecase - create new usercase
func NewUserUsecase(userRepo user.Repository) user.Usecase {
	return userUsecase{repo: userRepo}
}

// SelectUserByID - return user`s data by ID
func (u userUsecase) SelectUserByID(id int64) (*model.User, error) {
	return u.repo.SelectUserByID(id)
}

// SelectUserByUsername - return user`s data by username
func (u userUsecase) SelectUserByUsername(username string) (*model.User, error) {
	return u.repo.SelectUserByUsername(username)
}

// CreateUser - create new user
func (u userUsecase) CreateUser(user *model.User) (int64, error) {
	return u.repo.CreateUser(user)
}

// UpdateUser - update user`s data
func (u userUsecase) UpdateUser(user *model.User) (int64, error) {
	return u.repo.UpdateUser(user)
}

// DeleteUser - delete user data
func (u userUsecase) DeleteUser(id int64) (int64, error) {
	return u.repo.DeleteUser(id)
}

// SelectStoryByID - select story`s data
func (u userUsecase) SelectStoryByID(id int64) (*model.Story, error) {
	return u.repo.SelectStoryByID(id)
}

// SelectGenresByStoryID - select all genres of story
func (u userUsecase) SelectGenresByStoryID(id int64) ([]string, error) {
	return u.repo.SelectGenresByStoryID(id)
}

// SelectTopHeadingsStories - return top stories
func (u userUsecase) SelectTopHeadingsStories(headings map[string]string) (*model.StoryHeadings, error) {
	return u.repo.SelectTopHeadingsStories(headings)
}

// UpdateStoryViews - increment story views
func (u userUsecase) UpdateStoryViews(id int64) (int64, error) {
	return u.repo.UpdateStoryViews(id)
}

// UpdateStoryRating - update story rating
func (u userUsecase) UpdateStoryRating(reqRating *model.RequestRating) (int64, error) {
	return u.repo.UpdateStoryRating(reqRating)
}

// ListStories - return all stories
func (u userUsecase) ListStories() ([]*model.Story, error) {
	return u.repo.ListStories()
}

// CreateStory - create new story
func (u userUsecase) CreateStory(story *model.Story) (int64, error) {
	return u.repo.CreateStory(story)
}

// UpdateStory - update story record
func (u userUsecase) UpdateStory(story *model.Story) (int64, error) {
	return u.repo.UpdateStory(story)
}

// DeleteStory - delete story by ID
func (u userUsecase) DeleteStory(id int64) (int64, error) {
	return u.repo.DeleteStory(id)
}

// SelectView - return bool view
func (u userUsecase) SelectView(storyID int64, userID int64) (bool, error) {
	return u.repo.SelectView(storyID, userID)
}

// SelectRate - return previous rate and rating
func (u userUsecase) SelectRate(storyID int64, userID int64) (float64, bool, error) {
	return u.repo.SelectRate(storyID, userID)
}

// CreateView - create view in table storyRatingViews
func (u userUsecase) CreateView(elem *model.StoryRatingViews) (int64, error) {
	return u.repo.CreateView(elem)
}

// UpdateRating - set rating and store previous rating
func (u userUsecase) UpdateRating(elem *model.StoryRatingViews) (int64, error) {
	return u.repo.UpdateRating(elem)
}
