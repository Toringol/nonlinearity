package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Story - record structure for mongoDB
type Story struct {
	ID              bson.ObjectId `json:"id" bson:"_id"`
	Title           string        `json:"title" bson:"title"`
	Description     string        `json:"description" bson:"description"`
	Image           string        `json:"image" bson:"image"`
	StoryPath       string        `json:"storyPath" bson:"storyPath"`
	Author          string        `json:"author" bson:"author"`
	Genres          []string      `json:"genres" bson:"genres"`
	EditorChoice    bool          `json:"editorChoice" bson:"editorChoice"`
	Rating          float64       `json:"rating" bson:"rating"`
	Views           int           `json:"views" bson:"views"`
	PublicationDate time.Time     `json:"publicationDate" bson:"publicationDate"`
}
