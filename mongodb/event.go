package mongodb

import (
	"github.com/nikunicke/hiveboard"
)

type EventMongo struct {
	events []hiveboard.Event
}

var _ hiveboard.EventMongo = &EventMongo{}

func NewEventService() *EventMongo {
	return &EventMongo{
		events: nil,
	}
}

func (em *EventMongo) GetHBEvents(url string) ([]hiveboard.Event, error) {
	return nil, nil
}

// func (s *EventService) GetHBEvents(url string) ([]hiveboard.Event, error) {
// 	fmt.Println("Logging frim GetHBEvents")
// 	return nil, nil
// }

// func (s *EventService) GetEventByID(url string) (*hiveboard.Event, error) {
// 	return nil, nil
// }

// func (s *EventService) GetEventParticipants(url string) ([]hiveboard.Participant, error) {
// 	return nil, nil
// }

// func (s *EventService) GetUserEvents(url string) ([]hiveboard.Event, error) {
// 	return nil, nil
// }
