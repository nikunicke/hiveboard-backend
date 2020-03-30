package authorize

import (
	"log"
	"os"

	"github.com/nikunicke/hiveboard"
	"golang.org/x/oauth2"
)

const (
	ErrorState = hiveboard.Error("Invalid state value")
)

func GetURL() string {
	hiveboard.OauthConf = &oauth2.Config{
		ClientID:     os.Getenv("UID42"),
		ClientSecret: os.Getenv("SECRET42"),
		RedirectURL:  "http://localhost:3000/callback/",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api.intra.42.fr/oauth/authorize",
			TokenURL: "https://api.intra.42.fr/oauth/token",
		},
	}
	return hiveboard.OauthConf.AuthCodeURL(hiveboard.OauthState)
}

func GetToken(code string, state string) {
	var err error
	if state != hiveboard.OauthState {
		log.Fatal(ErrorState)
	}
	hiveboard.OauthToken, err = hiveboard.OauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatal(err)
	}
	hiveboard.Client = hiveboard.OauthConf.Client(oauth2.NoContext, hiveboard.OauthToken)
}
