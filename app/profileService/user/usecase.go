package user

import "github.com/Toringol/nonlinearity/app/profileService/model"

// Usecase - funcs interact with User
type Usecase interface {
	SelectUserByID(int64) (*model.User, error)
	CreateUser(*model.User) (int64, error)
	UpdateUser(*model.User) (int64, error)
	DeleteUSer(int64) (int64, error)
}
