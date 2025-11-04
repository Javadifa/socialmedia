package entity

import "time"

type Post struct {
	ID     int64
	UserID int64
	//Handle    string
	Caption   string
	ImageURL  string
	CreatedAt time.Time
}
