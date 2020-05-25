package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// StoryHeadings - headings structure
type StoryHeadings struct {
	Headings []*Heading `json:"headings"`
}

// Heading - some heading with top 10 stories
type Heading struct {
	Title   string   `json:"title"`
	Stories []*Story `json:"stories"`
}

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
	RatingsNumber   int           `json:"ratingsNumber" bson:"ratingsNumber"`
	Rating          float64       `json:"rating" bson:"rating"`
	Views           int           `json:"views" bson:"views"`
	PublicationDate time.Time     `json:"publicationDate" bson:"publicationDate"`
}

// RequestRating - request to change rating
type RequestRating struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	Rating float64       `json:"rating" bson:"rating"`
}

// RequestIDStory - request to change views
type RequestIDStory struct {
	ID bson.ObjectId `json:"id" bson:"_id"`
}
