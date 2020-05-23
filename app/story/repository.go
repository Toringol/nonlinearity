package story

import "github.com/Toringol/nonlinearity/app/model"

// Repository - funcs interact with DataBase
type Repository interface {
	SelectByID(string) (*model.Story, error)
	List() ([]*model.Story, error)
	Create(*model.Story) error
	Update(*model.Story) error
	Delete(string) error
}
