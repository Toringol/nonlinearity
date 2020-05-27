package model

import (
	"time"
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
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	StoryPath   string `json:"storyPath"`
	Author      string `json:"author"`
	// Genres          []string      `json:"genres" bson:"genres"`
	EditorChoice    bool      `json:"editorChoice"`
	RatingsNumber   int       `json:"ratingsNumber"`
	Rating          float64   `json:"rating"`
	Views           int       `json:"views"`
	PublicationDate time.Time `json:"publicationDate"`
}

// RequestRating - request to change rating
type RequestRating struct {
	ID     int64   `json:"id"`
	Rating float64 `json:"rating"`
}

// RequestIDStory - request to change views
type RequestIDStory struct {
	ID int64 `json:"id"`
}
