package repository

import (
	"database/sql"
	"log"

	"github.com/Toringol/nonlinearity/app/model"
	"github.com/Toringol/nonlinearity/app/user"
	"github.com/spf13/viper"
)

// Repository - Database implemetation
type Repository struct {
	DB *sql.DB
}

// NewUserMemoryRepository - create connection and return new repository
func NewUserMemoryRepository() user.Repository {
	dsn := viper.GetString("database")
	dsn += "&charset=utf8"
	dsn += "&interpolateParams=true"

	db, err := sql.Open("mysql", dsn)
	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		log.Println("Error while Ping")
	}

	return &Repository{
		DB: db,
	}
}

// ListStories - return all stories
func (repo *Repository) ListStories() ([]*model.Story, error) {
	stories := []*model.Story{}

	rows, err := repo.DB.Query("SELECT id, title, description, image, storyPath, author," +
		"editorChoice, rating, views, publicationDate FROM stories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		story := &model.Story{}
		err = rows.Scan(&story.ID, &story.Title, &story.Description, &story.Image, &story.StoryPath,
			&story.Author, &story.EditorChoice, &story.Rating, &story.Views, &story.PublicationDate)
		if err != nil {
			return nil, err
		}
		stories = append(stories, story)
	}

	return stories, nil
}

// SelectTopHeadingsStories - return top headings from story DB
func (repo *Repository) SelectTopHeadingsStories(headings map[string]string) (*model.StoryHeadings, error) {
	storyHeadings := &model.StoryHeadings{}

	for title, heading := range headings {
		stories := []*model.Story{}

		rows, err := repo.DB.Query("SELECT id, title, description, image, storyPath, author,"+
			"editorChoice, rating, views, publicationDate"+
			"FROM stories"+
			"ORDER BY ? DESC LIMIT 10",
			heading,
		)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			story := &model.Story{}
			err = rows.Scan(&story.ID, &story.Title, &story.Description, &story.Image, &story.StoryPath,
				&story.Author, &story.EditorChoice, &story.Rating, &story.Views, &story.PublicationDate)
			if err != nil {
				return nil, err
			}
			stories = append(stories, story)
		}

		headingRes := &model.Heading{
			Title:   title,
			Stories: stories,
		}

		storyHeadings.Headings = append(storyHeadings.Headings, headingRes)
	}

	return storyHeadings, nil
}

// SelectUserByID - select all user`s data by ID
func (repo *Repository) SelectUserByID(id int64) (*model.User, error) {
	record := &model.User{}
	err := repo.DB.
		QueryRow("SELECT id, login, password, avatar FROM users WHERE id = ?", id).
		Scan(&record.ID, &record.Username, &record.Password, &record.Avatar)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// SelectStoryByID - select all story`s data by ID
func (repo *Repository) SelectStoryByID(id int64) (*model.Story, error) {
	record := &model.Story{}
	err := repo.DB.
		QueryRow("SELECT id, title, description, image, storyPath, author,"+
			"editorChoice, rating, views, publicationDate FROM stories WHERE id = ?", id).
		Scan(&record.ID, &record.Title, &record.Description, &record.Image, &record.StoryPath,
			&record.Author, &record.EditorChoice, &record.Rating, &record.Views, &record.PublicationDate)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// SelectUserByUsername - select all user`s data by username
func (repo *Repository) SelectUserByUsername(username string) (*model.User, error) {
	record := &model.User{}
	err := repo.DB.
		QueryRow("SELECT id, login, password, avatar FROM users WHERE login = ?", username).
		Scan(&record.ID, &record.Username, &record.Password, &record.Avatar)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// CreateUser - create new User in dataBase with default avatar
func (repo *Repository) CreateUser(elem *model.User) (int64, error) {
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

// CreateStory - create new story in story table
func (repo *Repository) CreateStory(elem *model.Story) (int64, error) {
	result, err := repo.DB.Exec(
		"INSERT INTO stories (`title`, `description`, `image`, `storyPath`, `author`,"+
			"`editorChoice`, `rating`, `views`, `publicationDate`"+
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		elem.Title,
		elem.Description,
		elem.Image,
		elem.StoryPath,
		elem.Author,
		elem.EditorChoice,
		elem.Rating,
		elem.Views,
		elem.PublicationDate,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// UpdateUser - update user`s data in DataBase
func (repo *Repository) UpdateUser(elem *model.User) (int64, error) {
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

// UpdateStory - update story`s data in DataBase
func (repo *Repository) UpdateStory(elem *model.Story) (int64, error) {
	result, err := repo.DB.Exec(
		"UPDATE stories SET"+
			"`title` = ?"+
			",`description` = ?"+
			",`image` = ?"+
			",`storyPath` = ?"+
			",`author` = ?"+
			",`editorChoice` = ?"+
			",`publicationDate` = ?"+
			"WHERE id = ?",
		elem.Title,
		elem.Description,
		elem.Image,
		elem.StoryPath,
		elem.Author,
		elem.EditorChoice,
		elem.PublicationDate,
		elem.ID,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// UpdateStoryViews - increment story views
func (repo *Repository) UpdateStoryViews(id int64) (int64, error) {
	result, err := repo.DB.Exec(
		"UPDATE stories SET"+
			" `views` = `views` + 1"+
			" WHERE id = ?",
		id,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// UpdateStoryRating - recount story rating
func (repo *Repository) UpdateStoryRating(reqRating *model.RequestRating) (int64, error) {
	result, err := repo.DB.Exec(
		"UPDATE stories SET"+
			" `ratingsNumber` = `ratingsNumber` + 1"+
			" `rating` = (`rating` * (`ratingsNumber` - 1) + ?) / `ratingsNumber`"+
			" WHERE id = ?",
		reqRating.Rating,
		reqRating.ID,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// DeleteUser - delete user`s record in DataBase
func (repo *Repository) DeleteUser(id int64) (int64, error) {
	result, err := repo.DB.Exec(
		"DELETE FROM users WHERE id = ?",
		id,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

// DeleteStory - delete story record from DataBase
func (repo *Repository) DeleteStory(id int64) (int64, error) {
	result, err := repo.DB.Exec(
		"DELETE FROM stories WHERE id = ?",
		id,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
