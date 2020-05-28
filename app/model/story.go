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

// Story - record structure of Story
type Story struct {
	ID              int64     `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	Image           string    `json:"image"`
	StoryPath       string    `json:"storyPath"`
	Author          string    `json:"author"`
	EditorChoice    bool      `json:"editorChoice"`
	Genres          []string  `json:"genres"`
	RatingsNumber   int64     `json:"ratingsNumber"`
	Rating          float64   `json:"rating"`
	Views           int64     `json:"views"`
	PublicationDate time.Time `json:"publicationDate"`
}

// GenreTable - record structure of Genre
type GenreTable struct {
	StoryID int64  `json:"storyID"`
	Genre   string `json:"genre"`
}

// StoryRatingViews - table for rating and views
type StoryRatingViews struct {
	StoryID      int64   `json:"storyID"`
	UserID       int64   `json:"userID"`
	View         bool    `json:"view"`
	Rating       bool    `json:"rating"`
	PreviousRate float64 `json:"previousRating"`
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
