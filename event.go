package hiveboard

import (
	"time"
)

// EventService ...
type EventService interface {
	GetEvents(url string) ([]Event, error)
	GetEventByID(url string) (*Event, error)
	GetEventParticipants(url string) ([]Participant, error)
	GetUserEvents(url string) ([]Event, error)
}

// EventMongo ... This can be removed. We can use the same interface as above.
// Just need to add EventParticipant and UserEvents features for our own data
type EventMongo interface {
	GetEvents() ([]Event, error)
	GetEventByID(id string) (*Event, error)
	PostEvent(event Event) (*Event, error)
}

// EventService2 is meant to combine interfaces
type EventService2 struct {
	API42   EventService
	Mongodb EventMongo
}

// NewE creates a new EventService2
func NewE() *EventService2 {
	return &EventService2{
		API42:   nil,
		Mongodb: nil,
	}
}

type Wrapper struct {
	API42Error   string  `json:"api42error"`
	Mongo42Error string  `json:"mongo42error"`
	Data         []Event `json:"data"`
}

// Event ...
type Event struct {
	ID             interface{} `bson:"_id,omitempty" json:"id,integer"`
	Name           string      `bson:"name" json:"name" validate:"required,min=5,max=200"`
	Description    string      `bson:"description" json:"description" validate:"required,min=5,max=2000"`
	Location       string      `bson:"location" json:"location" validate:"required"`
	Kind           string      `bson:"kind" json:"kind"`
	MaxPeople      int         `bson:"max_people" json:"max_people"`
	NbrSubscribers int         `bson:"nbr_subscribers" json:"nbr_subscribers"`
	BeginAt        time.Time   `bson:"begin_at" json:"begin_at" validate:"required,gtfield=CreatedAt,ltfield=EndAt"`
	EndAt          time.Time   `bson:"end_at" json:"end_at" validate:"required,gtfield=BeginAt"`
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
	Hiveboard                 bool        `bson:"hiveboard" json:"hiveboard"`
}

// Participant ...
type Participant struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	URL   string `json:"url"`
}
