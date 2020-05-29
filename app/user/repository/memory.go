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

		rows, err := repo.DB.Query("SELECT id, title, image"+
			"editorChoice, rating"+
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

// SelectUserFavouritesByID - select user favourites
func (repo *Repository) SelectUserFavouritesByID(id int64) (*model.FavouriteCategories, error) {
	record := &model.FavouriteCategories{}

	err := repo.DB.
		QueryRow("SELECT drama, romance, comedy, horror, detective, fantasy, action, realism"+
			" FROM userFavourites WHERE id = ?", id).
		Scan(&record.Drama, &record.Romance, &record.Comedy, &record.Horror, &record.Detective,
			&record.Fantasy, &record.Action, &record.Realism)
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

// SelectView - check view of story
func (repo *Repository) SelectView(storyID int64, userID int64) (bool, error) {
	view := false

	err := repo.DB.
		QueryRow("SELECT view FROM storyRaringViews WHERE storyID = ? AND userID = ?", storyID, userID).
		Scan(&view)
	if err != nil {
		return false, err
	}

	return view, nil
}

// SelectRate - check rating of story
func (repo *Repository) SelectRate(storyID int64, userID int64) (float64, bool, error) {
	rating := false
	previousRate := float64(0)

	err := repo.DB.
		QueryRow("SELECT rating, previousRate FROM storyRaringViews WHERE storyID = ? AND userID = ?", storyID, userID).
		Scan(&rating, &previousRate)
	if err != nil {
		return 0, false, err
	}

	return previousRate, rating, nil
}

// SelectGenresByStoryID - return all genres of story
func (repo *Repository) SelectGenresByStoryID(id int64) ([]string, error) {
	genres := []string{}

	rows, err := repo.DB.Query("SELECT id, genre FROM genres WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		genre := &model.GenreTable{}

		err = rows.Scan(&genre.StoryID, &genre.Genre)
		if err != nil {
			return nil, err
		}

		genres = append(genres, genre.Genre)
	}

	return genres, nil
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

// CreateUserFavourites - create new record in DB about user Favourites
func (repo *Repository) CreateUserFavourites(id int64, elem *model.FavouriteCategories) (int64, error) {
	result, err := repo.DB.Exec(
		"INSERT INTO userFavourites (`id`, `drama`, `romance`, `comedy`, `horror`, `detective`,"+
			"`fantasy`, `action`, `realism`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		id,
		elem.Drama,
		elem.Romance,
		elem.Comedy,
		elem.Horror,
		elem.Detective,
		elem.Fantasy,
		elem.Action,
		elem.Realism,
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

// CreateView - create new view in storyRatingViews table
func (repo *Repository) CreateView(elem *model.StoryRatingViews) (int64, error) {
	result, err := repo.DB.Exec(
		"INSERT INTO storyRaringViews (`storyID`, `userID`, `view`) VALUES (?, ?, ?)",
		elem.StoryID,
		elem.UserID,
		elem.View,
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

// UpdateUserFavourites - increment favourite genres
func (repo *Repository) UpdateUserFavourites(id int64, genres *model.FavouriteCategories) (int64, error) {
	result, err := repo.DB.Exec(
		"UPDATE userFavourites SET"+
			" `drama` = `drama` + ?"+
			" `romance` = `romance` + ?"+
			" `comedy` = `comedy` + ?"+
			" `horror` = `horror` + ?"+
			" `detective` = `detective` + ?"+
			" `fantasy` = `fantasy` + ?"+
			" `action` = `action` + ?"+
			" `realism` = `realism` + ?"+
			" WHERE id = ?",
		genres.Drama,
		genres.Romance,
		genres.Comedy,
		genres.Horror,
		genres.Detective,
		genres.Fantasy,
		genres.Action,
		genres.Realism,
		id,
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

// UpdateRating - set rating on story
func (repo *Repository) UpdateRating(elem *model.StoryRatingViews) (int64, error) {
	elem.Rating = true
	result, err := repo.DB.Exec(
		"UPDATE storyRaringViews SET"+
			" `rating` = ?"+
			" `previousRate` = ?"+
			" WHERE storyID = ? AND userID = ?",
		elem.Rating,
		elem.PreviousRate,
		elem.StoryID,
		elem.UserID,
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
