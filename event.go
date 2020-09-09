package hiveboard

import (
	"encoding/json"
	"time"
)

type EventService interface {
	Get42Events(url string) ([]Event, error)
	// GetHBEvents(url string) ([]Event, error)
	GetEventByID(url string) (*Event, error)
	GetEventParticipants(url string) ([]Participant, error)
	GetUserEvents(url string) ([]Event, error)
}

type EventMongo interface {
	GetHBEvents(url string) ([]Event, error)
}

// lets try this

type EventService2 struct {
	API42   EventService
	Mongodb EventMongo
}

func NewE() *EventService2 {
	return &EventService2{
		API42:   nil,
		Mongodb: nil,
	}
}

// type EventService2 interface {
// 	EventMongo
// 	EventService
// }

type Event struct {
	ID             json.Number `bson:"_id,omitempty" json:"id"`
	Name           string      `bson:"name" json:"name"`
	Description    string      `bson:"description" json:"description"`
	Location       string      `bson:"location" json:"location"`
	Kind           string      `bson:"kind" json:"kind"`
	MaxPeople      int         `bson:"max_people" json:"max_people"`
	NbrSubscribers int         `bson:"nbr_subscribers" json:"nbr_subscribers"`
	BeginAt        time.Time   `bson:"begin_at" json:"begin_at"`
	EndAt          time.Time   `bson:"end_at" json:"end_at"`
	CampusIds      []int       `bson:"campus_ids" json:"campus_ids"`
	CursusIds      []int       `bson:"cursus_ids" json:"cursus_ids"`
	Themes         []struct {
		CreatedAt time.Time `bson:"created_at" json:"created_at"`
		ID        int       `bson:"id" json:"id"`
		Name      string    `bson:"name" json:"name"`
		UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	} `bson:"themes" json:"themes"`
	Waitlist                  interface{} `bson:"waitlist" json:"waitlist"`
	ProhibitionOfCancellation int         `bson:"prohibition_of_cancellation" json:"prohibition_of_cancellation"`
	CreatedAt                 time.Time   `bson:"created_at" json:"created_at"`
	UpdatedAt                 time.Time   `bson:"updated_at" json:"updated_at"`
	Tags                      []string    `bson:"tags" json:"tags"`
	Groups                    []string    `bson:"groups" json:"groups"`
}

type Participant struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	URL   string `json:"url"`
}
