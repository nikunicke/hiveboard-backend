package http

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/nikunicke/hiveboard"
)

const eventURL = "https://api.intra.42.fr/v2/events/"

type eventHandler struct {
	router       chi.Router
	baseURL      url.URL
	eventService hiveboard.EventService
}

func newEventHandler() *eventHandler {
	h := &eventHandler{router: chi.NewRouter()}
	h.router.Get("/", h.handleAllEvents)
	h.router.Get("/{id}", h.handleEventByID)
	h.router.Get("/{id}/users", h.handleEventParticipants)
	return h
}

func (h *eventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *eventHandler) handleAllEvents(w http.ResponseWriter, r *http.Request) {
	if hiveboard.Client == nil {
		http.Error(w, "Not Authorized", 401)
		return
	}

	events, err := h.eventService.GetEvents(eventURL)
	if err != nil {
		http.Error(w, "Internal Error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func (h *eventHandler) handleEventByID(w http.ResponseWriter, r *http.Request) {
	if hiveboard.Client == nil {
		http.Error(w, "Not Authorized", 401)
		return
	}
	eventID := chi.URLParam(r, "id")
	event, err := h.eventService.GetEventByID(eventURL + eventID)
	if err != nil {
		http.Error(w, "Internal Error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}

func (h *eventHandler) handleEventParticipants(w http.ResponseWriter, r *http.Request) {
	if hiveboard.Client == nil {
		http.Error(w, "Not Authorized", 401)
		return
	}
	eventID := chi.URLParam(r, "id")
	participants, err := h.eventService.GetEventParticipants(eventURL + eventID + "/users")
	if err != nil {
		http.Error(w, "Internal Error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(participants)
}
