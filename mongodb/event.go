package mongodb

import (
	"github.com/nikunicke/hiveboard"
)

type EventService struct {
	events []hiveboard.Event
}

// var _ hiveboard.EventService2 = &EventService{}

func NewEventService() *EventService {
	return &EventService{
		events: nil,
	}
}

func (s *EventService) GetEvents(url string) ([]hiveboard.Event, error) {
	return nil, nil
}

func (s *EventService) GetEventByID(url string) (*hiveboard.Event, error) {
	return nil, nil
}

func (s *EventService) GetEventParticipants(url string) ([]hiveboard.Participant, error) {
	return nil, nil
}

func (s *EventService) GetUserEvents(url string) ([]hiveboard.Event, error) {
	return nil, nil
}
