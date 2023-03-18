package config

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	State       = "state"
	RedirectURL = os.Getenv("REDIRECT_URL")

	GoogleClientID  = os.Getenv("GOOGLE_CLIENT_ID")
	GoogleClientSec = os.Getenv("GOOGLE_CLIENT_SECRET")
	GoogleScopes    = []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	}
)

func GetGoogleOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  RedirectURL,
		ClientID:     GoogleClientID,
		ClientSecret: GoogleClientSec,
		Scopes:       GoogleScopes,
		Endpoint:     google.Endpoint,
	}
}
