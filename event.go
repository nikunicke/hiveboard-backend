package hiveboard

import (
	"time"
)

// EventService ...
type EventService interface {
	GetEvents(url string) ([]Event, error)
	GetEventByID(url string) (*Event, error)
	GetEventUsers(url string) ([]EventUser, error)
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

// EventWrapper contains errors and the combined data from the 42API and our own
type EventWrapper struct {
	API42Error   string  `json:"api42error"`
	Mongo42Error string  `json:"mongo42error"`
	Data         []Event `json:"data"`
}

// Event represents an event
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

// EventUser represents a participant to an event
type EventUser struct {
	ID    interface{} `bson:"_id" json:"id"`
	Login string      `bson:"login" json:"login"`
	URL   string      `bson:"url" json:"url"`
}
