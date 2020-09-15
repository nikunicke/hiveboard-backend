package api42

import (
	"encoding/json"
	"io/ioutil"
	"time"

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
	beginAt := time.Now().Format("2006-01-02T15:04:05.000Z")
	endAt := time.Now().AddDate(1, 0, 0).Format("2006-01-02T15:04:05.000Z")
	response, err := hiveboard.Client.Get(url + "?range[begin_at]=" + beginAt + "," + endAt)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
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
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if err = json.Unmarshal(body, &event); err != nil {
		return nil, err
	}
	return event, nil
}

func (s *EventService) GetEventUsers(url string) ([]hiveboard.EventUser, error) {
	var participants []hiveboard.EventUser

	response, err := hiveboard.Client.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if err = json.Unmarshal(body, &participants); err != nil {
		return nil, err
	}
	return participants, nil
}

// Think about moving the below method to the userHandler. Right now the url path
// gets a liqttle bit messy --> domain/events/users/:user_id/events

// GetUserEvents domain/users/:user_id/events would work better
// check that ot works with the user bolt package
func (s *EventService) GetUserEvents(url string) ([]hiveboard.Event, error) {
	var events []hiveboard.Event
	response, err := hiveboard.Client.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if body[0] != '[' {
		return nil, hiveboard.UserNotFound
	}
	defer response.Body.Close()
	if err = json.Unmarshal(body, &events); err != nil {
		return nil, err
	}
	return events, nil
}
