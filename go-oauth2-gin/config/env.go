package config

import "os"

var (
	EnvGoogleClientID     = os.Getenv("GOOGLE_CLIENT_ID")
	EnvGoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	EnvGoogleRedirectURL  = os.Getenv("GOOGLE_REDIRECT_URL")
)
