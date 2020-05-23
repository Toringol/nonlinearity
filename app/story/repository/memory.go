package repository

import (
	"log"

	"github.com/Toringol/nonlinearity/app/model"
	"github.com/Toringol/nonlinearity/app/story"
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// StoryRepository - Story Database implemetation
type StoryRepository struct {
	DB *mgo.Database
}

// NewStoryMemoryRepository - create new story repository
func NewStoryMemoryRepository() story.Repository {
	dsn := viper.GetString("databaseStoriesConfig")

	session, err := mgo.Dial(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return &StoryRepository{
		DB: session.DB("storyDB"),
	}
}

// collection - return collection of story database
func (repo *StoryRepository) collection() *mgo.Collection {
	return repo.DB.C("stories")
}

// List - return all information about all stories in DB
func (repo *StoryRepository) List() ([]*model.Story, error) {
	stories := []*model.Story{}

	if err := repo.collection().Find(bson.M{}).All(&stories); err != nil {
		return nil, err
	}

	return stories, nil
}

// SelectByID - return story info by ID
func (repo *StoryRepository) SelectByID(id string) (*model.Story, error) {
	story := &model.Story{}

	if err := repo.collection().Find(bson.M{"_id": id}).One(&story); err != nil {
		return nil, err
	}

	return story, nil
}

// Create - create new story in story DB
func (repo *StoryRepository) Create(story *model.Story) error {
	story.ID = bson.NewObjectId()

	return repo.collection().Insert(story)
}

// Update - update records in story DB
func (repo *StoryRepository) Update(story *model.Story) error {
	oldStory := &model.Story{}

	if err := repo.collection().Find(bson.M{"_id": story.ID}).One(&oldStory); err != nil {
		return err
	}

	oldStory.Title = story.Title
	oldStory.Description = story.Description
	oldStory.Image = story.Image
	oldStory.StoryPath = story.StoryPath
	oldStory.Author = story.Author
	oldStory.EditorChoice = story.EditorChoice
	oldStory.Genres = story.Genres
	oldStory.PublicationDate = story.PublicationDate
	oldStory.Rating = story.Rating
	oldStory.Views = story.Views

	return repo.collection().Update(bson.M{"_id": oldStory.ID}, &oldStory)
}

// Delete - delete story from story DB
func (repo *StoryRepository) Delete(id string) error {
	return repo.collection().Remove(bson.M{"_id": id})
}
