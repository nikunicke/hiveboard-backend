package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nikunicke/hiveboard"
)

func GetUser(url string) (hiveboard.User, error) {
	var user hiveboard.User

	if hiveboard.OauthToken == nil {
		return user, hiveboard.ErrAuth
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return user, err
	}
	request.Header.Add("Authorization", "Bearer "+hiveboard.OauthToken.AccessToken)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return user, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return user, err
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}
