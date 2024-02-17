package main

import (
	"time" // Provides functionality for measuring and displaying time.

	"github.com/Gustavo-Villar/TideTracker/internal/database" // Internal package for database interactions.
	"github.com/google/uuid"                                  // Package for generating and handling UUIDs.
)

// User represents a user entity with identification, name, creation and update timestamps, and an API key for authentication.
type User struct {
	ID        uuid.UUID `json:"id"`         // Unique identifier for the user.
	Name      string    `json:"name"`       // Name of the user.
	CreatedAt time.Time `json:"created_at"` // Timestamp of when the user was created.
	UpdatedAt time.Time `json:"updated_at"` // Timestamp of the last update to the user's information.
	ApiKey    string    `json:"api_key"`    // API key for user authentication.
}

// databaseUserToUser converts a User entity from the database model to the application-level User struct.
func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey:    dbUser.ApiKey,
	}
}

// Feed represents an RSS feed entity with properties like ID, name, URL, associated user ID, and timestamps.
type Feed struct {
	ID        uuid.UUID `json:"id"`         // Unique identifier for the feed.
	Name      string    `json:"name"`       // Name of the feed.
	Url       string    `json:"url"`        // URL of the RSS feed.
	UserID    uuid.UUID `json:"user_id"`    // ID of the user who owns or follows the feed.
	CreatedAt time.Time `json:"created_at"` // Timestamp of when the feed was created.
	UpdatedAt time.Time `json:"updated_at"` // Timestamp of the last update to the feed.
}

// databaseFeedToFeed converts a Feed entity from the database model to the application-level Feed struct.
func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
	}
}

// databaseFeedsToFeeds converts a slice of Feed entities from the database model to a slice of application-level Feed structs.
func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}
	return feeds
}

// FeedFollow represents a relationship entity between a user and a feed, indicating a subscription.
type FeedFollow struct {
	ID        uuid.UUID `json:"id"`         // Unique identifier for the feed follow relationship.
	UserID    uuid.UUID `json:"user_id"`    // ID of the user following the feed.
	FeedID    uuid.UUID `json:"feed_id"`    // ID of the feed being followed.
	CreatedAt time.Time `json:"created_at"` // Timestamp of when the follow relationship was created.
	UpdatedAt time.Time `json:"updated_at"` // Timestamp of the last update to the follow relationship.
}

// databaseFeedFollowToFeedFollow converts a FeedFollow entity from the database model to the application-level FeedFollow struct.
func databaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
	}
}

// databaseFeedFollowsToFeedFollows converts a slice of FeedFollow entities from the database model to a slice of application-level FeedFollow structs.
func databaseFeedFollowsToFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}
	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, databaseFeedFollowToFeedFollow(dbFeedFollow))
	}
	return feedFollows
}

// Post represents a blog post or article entity within an RSS feed, including metadata like title, description, and publication date.
type Post struct {
	ID          uuid.UUID `json:"id"`           // Unique identifier for the post.
	Title       string    `json:"title"`        // Title of the post.
	Description *string   `json:"description"`  // Optional description or summary of the post.
	Url         string    `json:"url"`          // URL to the full post.
	FeedID      uuid.UUID `json:"feed_id"`      // ID of the feed this post belongs to.
	PublishedAt time.Time `json:"published_at"` // Timestamp of when the post was published.
	CreatedAt   time.Time `json:"created_at"`   // Timestamp of when the post was added to the database.
	UpdatedAt   time.Time `json:"updated_at"`   // Timestamp of the last update to the post.
}

// databasePostToPost converts a Post entity from the database model to the application-level Post struct.
func databasePostToPost(dbPost database.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		Title:       dbPost.Title,
		Description: description,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
		PublishedAt: dbPost.PublishedAt,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
	}
}

// databasePostsToPosts converts a slice of Post entities from the database model to a slice of application-level Post structs.
func databasePostsToPosts(dbPosts []database.Post) []Post {
	posts := []Post{}
	for _, dbPost := range dbPosts {
		posts = append(posts, databasePostToPost(dbPost))
	}
	return posts
}
