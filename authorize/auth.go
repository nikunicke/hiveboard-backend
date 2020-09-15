package authorize

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/nikunicke/hiveboard"
	"golang.org/x/oauth2"
)

// Exceptions
const (
	ErrorState = hiveboard.Error("Invalid state value")
)

// GetURL configures Oauth2 and returns the AuthCodeURL string
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

// GetToken communicates with the AuthURL and exchanges our code with a token
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
	// This should be saved somewhere. No need to get the full username and stuff...
	res, err := hiveboard.Client.Get("https://api.intra.42.fr/oauth/token/info")
	if err != nil {
		fmt.Println(err)
	}
	bod, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bod))
}
