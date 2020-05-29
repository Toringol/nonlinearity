package model

// User - user data in DataBase
type User struct {
	ID       int64  `json:"id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Avatar   string `json:"avatar"`
	// the best categories of User in slice byte format
	Favourited *FavouriteCategories `json:"favourited,omitempty"`
}

// FavouriteCategories - the number of counters
// preferences of Use
type FavouriteCategories struct {
	Drama     int64 `json:"drama"`
	Romance   int64 `json:"romance"`
	Comedy    int64 `json:"comedy"`
	Horror    int64 `json:"horror"`
	Detective int64 `json:"detective"`
	Fantasy   int64 `json:"fantasy"`
	Action    int64 `json:"action"`
	Realism   int64 `json:"realism"`
}
