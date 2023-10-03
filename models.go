package main

import (
	"rss/internal/database"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json: "id"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
	Name      string    `json: "name"`
	ApiKey    string    `json: "ApiKey`
}

// connection that takes a user from databse and returns a user struct
func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json: "id"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
	Name      string    `json: "name"`
	Url       string    `json: "url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}
<<<<<<< HEAD

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json: "id"`
	CreatedAt time.Time `json: "created_at"`
	UpdatedAt time.Time `json: "updated_at"`
	Name      string    `json: "name"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func databaseFeedFollowsToFeedFollows(dbFeedFollows database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollows.ID,
		CreatedAt: dbFeedFollows.CreatedAt,
		UpdatedAt: dbFeedFollows.UpdatedAt,
		UserID:    dbFeedFollows.UserID,
		FeedID:    dbFeedFollows.FeedsID,
	}
}
=======
>>>>>>> parent of 271b488 (created database Feeds To Feeds function in models.go)
