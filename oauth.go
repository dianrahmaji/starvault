package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
	"gorm.io/gorm"
)

var conf *oauth2.Config

type AuthResponse struct {
	Message string `json:"message"`
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	state := "state-token"

	conf = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes:       []string{youtube.YoutubeScope},
	}

	url := conf.AuthCodeURL(state, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleOAuthCallback(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	state := r.FormValue("state")

	if state != "state-token" {
		http.Error(w, "State token does not match", http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")

	if code == "" {
		http.Error(w, "Authorization code missing", http.StatusBadRequest)
		return
	}

	token, err := conf.Exchange(context.Background(), code)

	if err != nil {
		http.Error(w, "Failed to exchange code "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully logged in!\n")
	fmt.Fprintf(w, "Access Token: %s\n", token.AccessToken)
	fmt.Fprintf(w, "Refresh Token: %s\n", token.RefreshToken)

	oauthToken := OAuthToken{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}

	if err := db.Create(&oauthToken).Error; err != nil {
		http.Error(w, "Failed to save token", http.StatusInternalServerError)
		return
	}

	response := AuthResponse{
		Message: "success",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)

}
