package repository

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Toringol/nonlinearity/app/model"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestSelectByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var elemID int64 = 1

	rows := sqlmock.
		NewRows([]string{"id", "login", "password", "avatar"})
	expect := []*model.User{
		{elemID, "toringol", "12345", "default", model.UserPersonalData{}},
	}

	for _, item := range expect {
		rows = rows.AddRow(item.ID, item.Username, item.Password, item.Avatar)
	}

	mock.
		ExpectQuery("SELECT id, login, password, avatar FROM users WHERE").
		WithArgs(elemID).
		WillReturnRows(rows)

	repo := &UserRepository{
		DB: db,
	}

	item, err := repo.SelectByID(elemID)

	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], item)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT id, login, password, avatar FROM users WHERE").
		WithArgs(elemID).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.SelectByID(elemID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	// row scan error
	rows = sqlmock.NewRows([]string{"id", "login"}).
		AddRow(1, "username")

	mock.
		ExpectQuery("SELECT id, login, password, avatar FROM users WHERE").
		WithArgs(elemID).
		WillReturnRows(rows)

	_, err = repo.SelectByID(elemID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestSelectByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var elemID int64 = 1
	login := "toringol"

	rows := sqlmock.
		NewRows([]string{"id", "login", "password", "avatar"})
	expect := []*model.User{
		{elemID, "toringol", "12345", "default", model.UserPersonalData{}},
	}

	for _, item := range expect {
		rows = rows.AddRow(item.ID, item.Username, item.Password, item.Avatar)
	}

	mock.
		ExpectQuery("SELECT id, login, password, avatar FROM users WHERE").
		WithArgs(login).
		WillReturnRows(rows)

	repo := &UserRepository{
		DB: db,
	}

	item, err := repo.SelectByUsername(login)

	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(item, expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], item)
		return
	}

	// query error
	mock.
		ExpectQuery("SELECT id, login, password, avatar FROM users WHERE").
		WithArgs(login).
		WillReturnError(fmt.Errorf("db_error"))

	_, err = repo.SelectByUsername(login)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	// row scan error
	rows = sqlmock.NewRows([]string{"id", "login"}).
		AddRow(1, "toringol")

	mock.
		ExpectQuery("SELECT id, login, password, avatar FROM users WHERE").
		WithArgs(login).
		WillReturnRows(rows)

	_, err = repo.SelectByUsername(login)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &UserRepository{
		DB: db,
	}

	username := "login"
	password := "password"
	defaultAvatar := "/images/avatar.png"
	testItem := &model.User{
		Username: username,
		Password: password,
		Avatar:   defaultAvatar,
	}

	//ok query
	mock.
		ExpectExec(`INSERT INTO users`).
		WithArgs(username, password, defaultAvatar).
		WillReturnResult(sqlmock.NewResult(1, 1))

	id, err := repo.Create(testItem)
	if id != 1 {
		t.Errorf("bad id: want %v, have %v", id, 1)
		return
	}

	// query error
	mock.
		ExpectExec(`INSERT INTO users`).
		WithArgs(username, password, defaultAvatar).
		WillReturnError(fmt.Errorf("bad query"))

	_, err = repo.Create(testItem)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	// result error
	mock.
		ExpectExec(`INSERT INTO users`).
		WithArgs(username, password, defaultAvatar).
		WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("bad_result")))

	_, err = repo.Create(testItem)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var elemID int64 = 1

	rows := sqlmock.
		NewRows([]string{"id", "login", "password", "avatar"})
	testInput := []*model.User{
		{elemID, "toringol", "12345", "default", model.UserPersonalData{}},
		{elemID + 1, "user", "pass", "default", model.UserPersonalData{}},
	}

	expect := &model.User{elemID, "sergey", "23623", "default", model.UserPersonalData{}}

	for _, item := range testInput {
		rows = rows.AddRow(item.ID, item.Username, item.Password, item.Avatar)
	}

	mock.
		ExpectExec(`UPDATE users SET`).
		WithArgs(expect.Username, expect.Password, expect.Avatar, expect.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := &UserRepository{
		DB: db,
	}

	rowsAffected, err := repo.Update(expect)
	if rowsAffected != 1 {
		t.Errorf("bad rowsAffected: want %v, have %v", rowsAffected, 1)
		return
	}

	// query error
	mock.
		ExpectExec(`UPDATE users SET`).
		WithArgs(expect.Username, expect.Password, expect.Avatar, expect.ID).
		WillReturnError(fmt.Errorf("bad query"))

	_, err = repo.Update(expect)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}

	// result error
	mock.
		ExpectExec(`UPDATE users SET`).
		WithArgs(expect.Username, expect.Password, expect.Avatar, expect.ID).
		WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("bad_result")))

	_, err = repo.Update(expect)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var elemID int64 = 1

	rows := sqlmock.
		NewRows([]string{"id", "login", "password", "avatar"})
	testInput := []*model.User{
		{elemID, "toringol", "12345", "default", model.UserPersonalData{}},
		{elemID + 1, "user", "pass", "default", model.UserPersonalData{}},
	}

	for _, item := range testInput {
		rows = rows.AddRow(item.ID, item.Username, item.Password, item.Avatar)
	}

	mock.
		ExpectExec(`DELETE FROM users WHERE`).
		WithArgs(elemID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := &UserRepository{
		DB: db,
	}

	rowsAffected, err := repo.Delete(elemID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if rowsAffected != 1 {
		t.Errorf("bad rowsAffected: want %v, have %v", rowsAffected, 1)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// query error
	mock.
		ExpectExec(`DELETE FROM users WHERE`).
		WithArgs(elemID).
		WillReturnError(fmt.Errorf("bad query"))

	_, err = repo.Delete(elemID)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// result error
	mock.
		ExpectExec(`DELETE FROM users WHERE`).
		WithArgs(elemID).
		WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("bad_result")))

	_, err = repo.Delete(elemID)
	if err == nil {
		t.Errorf("expected error, got nil")
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
