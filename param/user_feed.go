package param

import (
	"time"
)

type FeedRequest struct {
	UserID uint `json:"user_id"`
}

type FeedResponsePost struct {
	ID     int64
	UserID int64
	//Handle    string
	ImageURL  string
	Caption   string
	CreatedAt time.Time
}

type FeedResponse struct {
	Posts []FeedResponsePost
}
