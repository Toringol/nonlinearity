package model

// Session - data in redis Storage sessID - useraname, useragent
type Session struct {
	Username  string
	Useragent string
}

// SessionID - sessionID in redis Storage
type SessionID struct {
	ID string
}
