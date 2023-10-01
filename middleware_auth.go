package main

import (
	"fmt"
	"net/http"
	"rss/auth"
	"rss/internal/database"
)

// New authHandler type
// Same as http handler but with authed user as 3rd parameter
type authHandler func(http.ResponseWriter, *http.Request, database.User)

// Middle ware function to fix the miss match between auth handler and http handler
func (cfg *apiConfig) middleWareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Error getting api key: %s", err))
			return
		}
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Couldn't get user: %s", err))
			return
		}
		handler(w, r, user)
	}
}
