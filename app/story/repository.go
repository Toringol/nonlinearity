package story

import "github.com/Toringol/nonlinearity/app/model"

// Repository - funcs interact with DataBase
type Repository interface {
	SelectByID(string) (*model.Story, error)
	SelectTopHeadings(map[string]string) (*model.StoryHeadings, error)
	UpdateViews(string) error
	UpdateRating(*model.RequestRating) (*model.Story, error)
	List() ([]*model.Story, error)
	Create(*model.Story) error
	Update(*model.Story) error
	Delete(string) error
}
