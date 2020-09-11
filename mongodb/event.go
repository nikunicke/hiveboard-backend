package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/nikunicke/hiveboard"
	"go.mongodb.org/mongo-driver/bson"
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
	return s.db.findAll("events")
}

// GetEventByID returns a given event based on ID
func (s *EventService) GetEventByID(id string) (*hiveboard.Event, error) {
	return s.db.findByID("events", id)
}

// func (s *EventService) PostEvent() error {

// }

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

// FindAll ...
func (db *MongoDB) findAll(collection string) ([]hiveboard.Event, error) {
	var results []hiveboard.Event

	fromDate := time.Now()
	cursor, err := db.db.Collection(collection).Find(context.TODO(), bson.M{"begin_at": bson.M{"$gt": fromDate}})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}

// findByID ...
func (db *MongoDB) findByID(collection string, id string) (*hiveboard.Event, error) {
	var event *hiveboard.Event

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := db.db.Collection(collection).FindOne(context.Background(), bson.M{"_id": objectID})
	result.Decode(&event)
	return event, nil
}
