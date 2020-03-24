package main

import (
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
	uid := "044c5b79a54831c1cd1d65cf3a9e08b4d12f885f013b86e86137c03dc478ca7b"
	secret := "4764c22900f4deb6c07b78283db4ac143d5300fa5b1844a5bd8270b1119d3bd8"

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
	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}
	return token, nil
}
