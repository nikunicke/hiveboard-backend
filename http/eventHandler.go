package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nikunicke/hiveboard"
)

func GetEvents(url string) ([]hiveboard.Event, error) {
	var events []hiveboard.Event

	if hiveboard.OauthToken == nil {
		return nil, hiveboard.ErrAuth
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", "Bearer "+hiveboard.OauthToken.AccessToken)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}
