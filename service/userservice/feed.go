package userservice

import (
	"fmt"
	"log"
	"time"

	//"github.com/javadifa/socialmedia/entity"
	"github.com/javadifa/socialmedia/param"
)

func (s Service) GetUserFeed(req param.FeedRequest) (param.FeedResponse, error) {
	start := time.Now()
	posts, err := s.repo.GetUserFeedByID(req.UserID)
	if err != nil {
		log.Printf("Error fetching feed for user %d: %v", req.UserID, err)
		return param.FeedResponse{}, fmt.Errorf("couldn't fetch the feed: %w", err)
	}
	log.Printf("Fetched feed for user %d in %v", req.UserID, time.Since(start))

	responsePosts := make([]param.FeedResponsePost, len(posts))
	for i, p := range posts {
		responsePosts[i] = param.FeedResponsePost{
			ID:     p.ID,
			UserID: p.UserID,
			//Handle:    p.Handle,
			ImageURL:  p.ImageURL,
			Caption:   p.Caption,
			CreatedAt: p.CreatedAt,
		}
	}
	//currently no need of handle, but later based on join query and user table, it's possible to return the handle as well
	return param.FeedResponse{Posts: responsePosts}, nil
}
