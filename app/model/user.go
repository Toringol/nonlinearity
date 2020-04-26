package model

import (
	"encoding/json"
	"time"
)

// User - user data in DataBase
type User struct {
	ID               int64            `json:"id,omitempty"`
	Username         string           `json:"username"`
	Password         string           `json:"password,omitempty"`
	Avatar           string           `json:"avatar"`
	UserPersonalData UserPersonalData `json:"personalData"`
}

// UserPersonalData - personal data of User
type UserPersonalData struct {
	DateOfBirth  time.Time `json:"birthday"`
	Relationship string    `json:"relationship"`
	Status       string    `json:"status"`
}

// Favourites - the best categories of User
type Favourites struct {
	ID         int64           `json:"id"`
	UserID     int64           `json:"userid"`
	Favourited json.RawMessage `json:"favourited"`
}
