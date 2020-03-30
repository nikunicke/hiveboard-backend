package http

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/nikunicke/hiveboard"
)

// type EventService struct {
// 	events []hiveboard.Event
// }

// var _ hiveboard.EventService = &EventService{}

// func NewEventService() *EventService {
// 	return &EventService{
// 		events: nil,
// 	}
// }

type eventHandler struct {
	router       chi.Router
	baseURL      url.URL
	eventService hiveboard.EventService
}

func newEventHandler() *eventHandler {
	h := &eventHandler{router: chi.NewRouter()}
	h.router.Get("/", h.handleGet)
	// h.router.Get("/:id/", )
	return h
}

func (h *eventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *eventHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	if hiveboard.Client == nil {
		http.Error(w, "Not Authorized", 401)
		return
	}

	events, err := h.eventService.GetEvents("https://api.intra.42.fr/v2/events")
	if err != nil {
		http.Error(w, "Internal Error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

// func (s *EventService) GetEvents(url string) ([]hiveboard.Event, error) {
// 	var events []hiveboard.Event
// 	response, err := hiveboard.Client.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	body, err := ioutil.ReadAll((response.Body))
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Printf("%s\n", string([]byte(body)))
// 	return events, nil
// }

// func (s *EventService) GetEvents(url string) ([]hiveboard.Event, error) {
// 	var events []hiveboard.Event

// 	request, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	request.Header.Add("Authorization", "Bearer "+hiveboard.OauthToken.AccessToken)
// 	client := &http.Client{}
// 	response, err := client.Do(request)
// 	if err != nil {
// 		return nil, err
// 	}
// 	body, err := ioutil.ReadAll((response.Body))
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = json.Unmarshal(body, &events)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return events, nil
// }

// func (s *EventService) GetEventByID(url string) (*hiveboard.Event, error) {
// 	var event hiveboard.Event

// 	request, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	request.Header.Add("Authorization", "Bearer "+hiveboard.OauthToken.TokenType)
// 	client := &http.Client{}
// 	response, err := client.Do(request)
// 	if err != nil {
// 		return nil, err
// 	}
// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = json.Unmarshal(body, &event)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &event, nil
// }

// func GetEvents(url string) ([]hiveboard.Event, error) {
// 	var events []hiveboard.Event

// 	if hiveboard.OauthToken == nil {
// 		return nil, hiveboard.ErrAuth
// 	}
// 	request, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	request.Header.Add("Authorization", "Bearer "+hiveboard.OauthToken.AccessToken)
// 	client := &http.Client{}
// 	response, err := client.Do(request)
// 	if err != nil {
// 		return nil, err
// 	}
// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	err = json.Unmarshal(body, &events)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return events, nil
// }
