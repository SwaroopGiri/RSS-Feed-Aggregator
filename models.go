package main

import "github.com/swaroop-giri/GoAgg/internal/database"

func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey:    dbUser.ApiKey,
	}
}

func databaseFeedtoFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func databaseFeedstoFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbfeed := range dbFeeds {
		feeds = append(feeds, databaseFeedtoFeed(dbfeed))
	}
	return feeds
}

func databaseFeedFollowtoFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
	}
}

func databaseFeedFollowstoFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}
	for _, dbfeedfollow := range dbFeedFollows {
		feedFollows = append(feedFollows, databaseFeedFollowtoFeedFollow(dbfeedfollow))
	}
	return feedFollows
}

func databasePosttoPost(dbPost database.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}
	return Post{
		ID:          dbPost.ID,
		Title:       dbPost.Title,
		Url:         dbPost.Url,
		Description: description,
		FeedID:      dbPost.FeedID,
		PublishedAt: dbPost.PublishedAt,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
	}
}

func databasePoststoPosts(dbPosts []database.Post) []Post {
	posts := []Post{}
	for _, dbpost := range dbPosts {
		posts = append(posts, databasePosttoPost(dbpost))
	}
	return posts
}
