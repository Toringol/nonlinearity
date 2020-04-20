package model

// Story - record structure for mongoDB
type Story struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Type        string `json:"type"`
	Views       string `json:"views"`
	Rating      int    `json:"rating"`
}
