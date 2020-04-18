package repository

import (
	"database/sql"
	"log"

	"github.com/Toringol/nonlinearity/app/profileService/model"
	"github.com/Toringol/nonlinearity/app/profileService/user"
)

// NewUserMemoryRepository - create connection and return new repository
func NewUserMemoryRepository() user.Repository {
	dsn := "" // TODO: get dsn from cfg
	dsn += "&charset=utf8"
	dsn += "&interpolateParams=true"

	db, err := sql.Open("mysql", dsn)
	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		log.Println("Error while Ping")
	}

	return &UserRepository{
		DB: db,
	}
}

// UserRepository - Database implemetation
type UserRepository struct {
	DB *sql.DB
}

// SelectByID - select all user`s data by ID
func (repo *UserRepository) SelectByID(id int64) (*model.User, error) {
	record := &model.User{}
	err := repo.DB.
		QueryRow("SELECT id, login, password, avatar FROM users WHERE id = ?", id).
		Scan(&record.ID, &record.Username, &record.Password, &record.Avatar)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// Create - create new User in dataBase with default avatar
func (repo *UserRepository) Create(elem *model.User) (int64, error) {
	result, err := repo.DB.Exec(
		"INSERT INTO users (`login`, `password`, `avatar`) VALUES (?, ?, ?)",
		elem.Username,
		elem.Password,
		elem.Avatar,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// Update - update user`s data in DataBase
func (repo *UserRepository) Update(elem *model.User) (int64, error) {
	result, err := repo.DB.Exec(
		"UPDATE users SET"+
			"`login` = ?"+
			",`password` = ?"+
			",`avatar` = ?"+
			"WHERE id = ?",
		elem.Username,
		elem.Password,
		elem.Avatar,
		elem.ID,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// Delete - delete user`s record in DataBase
func (repo *UserRepository) Delete(id int64) (int64, error) {
	result, err := repo.DB.Exec(
		"DELETE FROM users WHERE id = ?",
		id,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
