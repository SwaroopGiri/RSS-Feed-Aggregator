package main

import (
	"fmt"
	"net/http"

	"github.com/swaroop-giri/GoAgg/internal/auth"
	"github.com/swaroop-giri/GoAgg/internal/database"
)

type authhandler func(http.ResponseWriter, *http.Request, database.User)

func (apicfg *apiConfig) authMiddleware(next authhandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondwithError(w, 401, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := apicfg.DB.GetUserByAPIKey(r.Context(), apikey)
		if err != nil {
			respondwithError(w, 400, fmt.Sprintf("Error getting user: %v", err))
			return
		}
		next(w, r, user)
	}
}
