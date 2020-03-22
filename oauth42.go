package main

import (
	"context"
	"errors"

	"golang.org/x/oauth2"
)

var (
	oauthConf *oauth2.Config
	// add State to ENV
	oauthState = "SuperRandomString"
)

func getAuthorization() string {
	// add uid and secret to ENV variables
	uid := YOUR_CLIENT_ID
	secret := YOUR_CLIENT_SECRET

	oauthConf = &oauth2.Config{
		ClientID:     uid,
		ClientSecret: secret,
		RedirectURL:  "http://localhost:3000/callback/",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.intra.42.fr/oauth/authorize",
			TokenURL: "https://api.intra.42.fr/oauth/token",
		},
	}
	url := oauthConf.AuthCodeURL(oauthState)
	return url
}

func getAccessToken(code string, state string) (*oauth2.Token, error) {
	if state != oauthState {
		return nil, errors.New("Invalid state value")
	}
	token, err := oauthConf.Exchange(context.TODO(), code)
	if err != nil {
		return nil, err
	}
	return token, nil
}
