package mongodb

import (
	"github.com/nikunicke/hiveboard"
)

// EventService represents a service to manage mongodb events
type EventService struct {
	db *MongoDB
}

var _ hiveboard.EventMongo = &EventService{}

// NewEventService returns a new instance of EventService
func NewEventService(db *MongoDB) *EventService {
	return &EventService{
		db: db,
	}
}

// GetEvents returns all events from the 'events' collection
func (s *EventService) GetEvents() ([]hiveboard.Event, error) {
	return s.db.FindAll("events")
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
