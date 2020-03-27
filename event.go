package hiveboard

import "time"

type EventService interface {
	GetEvents() (*[]Event, error)
	FindEventByID(id int) (*Event, error)
	GetEventParticipands(e *Event) (*[]Participant, error)
}

type Event struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Location       string    `json:"location"`
	Kind           string    `json:"kind"`
	MaxPeople      int       `json:"max_people"`
	NbrSubscribers int       `json:"nbr_subscribers"`
	BeginAt        time.Time `json:"begin_at"`
	EndAt          time.Time `json:"end_at"`
	CampusIds      []int     `json:"campus_ids"`
	CursusIds      []int     `json:"cursus_ids"`
	Themes         []struct {
		CreatedAt time.Time `json:"created_at"`
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"themes"`
	Waitlist                  interface{} `json:"waitlist"`
	ProhibitionOfCancellation int         `json:"prohibition_of_cancellation"`
	CreatedAt                 time.Time   `json:"created_at"`
	UpdatedAt                 time.Time   `json:"updated_at"`
}

type Participant struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	URL   string `json:"url"`
}
