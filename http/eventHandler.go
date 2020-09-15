package http

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/nikunicke/hiveboard"
)

const eventURL = "https://api.intra.42.fr/v2/"

type eventHandler struct {
	router        chi.Router
	baseURL       url.URL
	eventService  hiveboard.EventService
	eventService2 hiveboard.EventService2
}

func newEventHandler() *eventHandler {
	h := &eventHandler{router: chi.NewRouter()}
	h.router.Get("/", h.getAll)
	h.router.Get("/{eventID}", h.getEventByID)
	h.router.Get("/{eventID}/users", h.handleGetEventUsers)
	h.router.Get("/users/{userID}", h.handleGetUserEvents)
	// h.router.Get("/eventsusers", h.getEventsUsers)
	h.router.Post("/", h.postEvent)
	return h
}

func (h *eventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *eventHandler) getAll(w http.ResponseWriter, r *http.Request) {
	var events hiveboard.EventWrapper

	if hiveboard.Client == nil {
		http.Error(w, "Not Authorized", 401)
		return
	}
	API42Events, err := h.eventService2.API42.GetEvents("https://api.intra.42.fr/v2/campus/13/" + "events")
	if err != nil {
		events.API42Error = err.Error()
	}
	hiveboardEvents, err := h.eventService2.Mongodb.GetEvents()
	if err != nil {
		events.Mongo42Error = err.Error()
	}
	if events.API42Error != "" && events.Mongo42Error != "" {
		http.Error(w, "Internal server error", 500)
	}
	events.Data = append(API42Events, hiveboardEvents...)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func (h *eventHandler) getEventByID(w http.ResponseWriter, r *http.Request) {
	var err error
	var event *hiveboard.Event

	if hiveboard.Client == nil {
		http.Error(w, "Not Authorized", 401)
		return
	}
	eventID := chi.URLParam(r, "eventID")
	if isNumeric(eventID) {
		event, err = h.eventService2.API42.GetEventByID(eventURL + "events/" + eventID)
	} else {
		event, err = h.eventService2.Mongodb.GetEventByID(eventID)
	}
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}
	if event == nil {
		http.Error(w, "Event not found", 404)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func (h *eventHandler) handleGetEventUsers(w http.ResponseWriter, r *http.Request) {
	var err error
	var eventUsers []hiveboard.EventUser
	if hiveboard.Client == nil {
		http.Error(w, "Not Authorized", 401)
		return
	}
	eventID := chi.URLParam(r, "eventID")
	if isNumeric(eventID) {
		suffix := "events/" + eventID + "/users"
		eventUsers, err = h.eventService2.API42.GetEventUsers(eventURL + suffix)
	} else {
		http.Error(w, "Our API does not support this feature yet", 500)
	}
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(eventUsers)
}

func (h *eventHandler) handleGetUserEvents(w http.ResponseWriter, r *http.Request) {
	if hiveboard.Client == nil {
		http.Error(w, "Not Authorized", 401)
		return
	}
	userID := chi.URLParam(r, "userID") + "/"
	events, err := h.eventService.GetUserEvents(eventURL + "users/" + userID + "events")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func (h *eventHandler) postEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent hiveboard.Event
	if hiveboard.Client == nil {
		http.Error(w, "Not Authorized", 401)
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newEvent)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	res, err := h.eventService2.Mongodb.PostEvent(newEvent)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(res)
}

// func (h *eventHandler) handleSub(w http.ResponseWriter, r *http.Request) {
// 	// eventID := chi.URLParam(r, "eventID") + "/"
// 	// the below method might actually work, but we do not have permissions

// 	res, err := hiveboard.Client.PostForm(eventURL+"events_users", url.Values{"events_user[event_id]": {"4376"}, "events_user[user_id]": {"59634"}})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(string(body))
// 	fmt.Fprintf(w, "Subscribed")
// }

// 59634 me
// 4376 Q&A
// 59333 random

// func (h *eventHandler) getEventsUsers(w http.ResponseWriter, r *http.Request) {
// 	if hiveboard.Client == nil {
// 		http.Error(w, "Not Authorized", 401)
// 		return
// 	}
// 	response, err := hiveboard.Client.Get(eventURL + "events_users/?user_id=59634")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	body, err := ioutil.ReadAll(response.Body)
// 	fmt.Fprintf(w, "%s\n", string(body))
// }

func isNumeric(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}
