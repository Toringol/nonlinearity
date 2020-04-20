package user

import "github.com/Toringol/nonlinearity/app/model"

// Repository - funcs interact with DataBase
type Repository interface {
	SelectByID(int64) (*model.User, error)
	SelectByUsername(string) (*model.User, error)
	Create(*model.User) (int64, error)
	Update(*model.User) (int64, error)
	Delete(int64) (int64, error)
}
