package usecase

import (
	"github.com/Toringol/nonlinearity/app/profileService/model"
	"github.com/Toringol/nonlinearity/app/profileService/user"
)

// NewUserUsecase - create new usercase
func NewUserUsecase(userRepo user.Repository) user.Usecase {
	return userUsecase{repo: userRepo}
}

// userUsecase - connector database and user
type userUsecase struct {
	repo user.Repository
}

// SelectUserByID - return user`s data by ID
func (u userUsecase) SelectUserByID(id int64) (*model.User, error) {
	return u.repo.SelectByID(id)
}

// CreateUser - create new user
func (u userUsecase) CreateUser(user *model.User) (int64, error) {
	return u.repo.Create(user)
}

// UpdateUser - update user`s data
func (u userUsecase) UpdateUser(user *model.User) (int64, error) {
	return u.repo.Update(user)
}

// DeleteUser - delete user data
func (u userUsecase) DeleteUser(id int64) (int64, error) {
	return u.repo.Delete(id)
}
