package bolt

import (
	"encoding/json"
	"io/ioutil"

	"github.com/nikunicke/hiveboard"
)

type EventService struct {
	events []hiveboard.Event
}

var _ hiveboard.EventService = &EventService{}

func NewEventService() *EventService {
	return &EventService{
		events: nil,
	}
}

func (s *EventService) GetEvents(url string) ([]hiveboard.Event, error) {
	var events []hiveboard.Event
	response, err := hiveboard.Client.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll((response.Body))
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &events); err != nil {
		return nil, err
	}
	return events, nil
}

func (s *EventService) GetEventByID(url string) (*hiveboard.Event, error) {
	var event *hiveboard.Event

	response, err := hiveboard.Client.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err = json.Unmarshal(body, &event); err != nil {
		return nil, err
	}
	return event, nil
}

func (s *EventService) GetEventParticipants(url string) ([]hiveboard.Participant, error) {
	var participants []hiveboard.Participant

	response, err := hiveboard.Client.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &participants); err != nil {
		return nil, err
	}
	return participants, nil
}

// func (s *EventService) GetEvents(url string) ([]hiveboard.Event, error) {
// 	var events []hiveboard.Event

// 	request, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	request.Header.Add("Authorization", "Bearer "+hiveboard.OauthToken.AccessToken)
// 	client := &http.Client{}
// 	response, err := client.Do(request)
// 	if err != nil {
// 		return nil, err
// 	}
// 	body, err := ioutil.ReadAll((response.Body))
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = json.Unmarshal(body, &events)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return events, nil
// }
