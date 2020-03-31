package hiveboard

import "time"

type EventService interface {
	GetEvents(url string) ([]Event, error)
	GetEventByID(url string) (*Event, error)
	GetEventParticipants(url string) ([]Participant, error)
	GetUserEvents(url string) ([]Event, error)
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
	Tags                      []string    `json:"tags"`
	Groups                    []string    `json:"groups"`
}

type Participant struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	URL   string `json:"url"`
}
