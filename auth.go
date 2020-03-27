package hiveboard

import (
	"os"

	"golang.org/x/oauth2"
)

var (
	OauthConf  *oauth2.Config
	OauthState = os.Getenv("STATE42")
	OauthToken *oauth2.Token
)
