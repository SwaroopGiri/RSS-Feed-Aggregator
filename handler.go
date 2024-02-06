package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/swaroop-giri/GoAgg/internal/database"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondwithJSON(w, 200, map[string]string{"status": "ok"})
}

func handlererr(w http.ResponseWriter, r *http.Request) {
	respondwithError(w, 400, "Bad Request")
}

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	params := apiparameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error parsing JSON request body: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	respondwithJSON(w, 201, databaseUsertoUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondwithJSON(w, 200, databaseUsertoUser(user))
}

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	decoder := json.NewDecoder(r.Body)
	params := Feed{}
	err := decoder.Decode(&params)
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error parsing JSON request body: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error creating feed: %v", err))
		return
	}

	respondwithJSON(w, 201, databaseFeedtoFeed(feed))
}

func (apiCfg *apiConfig) HandlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error getting feeds: %v", err))
		return
	}
	respondwithJSON(w, 200, databaseFeedstoFeeds(feeds))
}

func (apiCfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)
	params := FeedFollow{}
	err := decoder.Decode(&params)
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error parsing JSON request body: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error creating feed follow: %v", err))
		return
	}

	respondwithJSON(w, 201, databaseFeedFollowtoFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollow(r.Context(), user.ID)
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error getting feed follows: %v", err))
		return
	}
	respondwithJSON(w, 200, databaseFeedFollowstoFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedfollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error parsing feed follow ID: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error deleting feed follow: %v", err))
		return
	}
	respondwithJSON(w, 200, map[string]string{"status": "unfollowed"})
}

func (apiCfg *apiConfig) handlerGetPostsforUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		respondwithError(w, 400, fmt.Sprintf("Error getting posts for User: %v", err))
		return
	}
	respondwithJSON(w, 200, databasePoststoPosts(posts))
}
