package http

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/nikunicke/hiveboard"
)

const userURL = "https://api.intra.42.fr/v2/"

type userHandler struct {
	router      chi.Router
	baseURL     url.URL
	userService hiveboard.UserService
}

func newUserHandler() *userHandler {
	h := &userHandler{router: chi.NewRouter()}
	h.router.Get("/", h.handleGetUser)
	return h
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h *userHandler) handleGetUser(w http.ResponseWriter, r *http.Request) {
	if hiveboard.Client == nil {
		http.Error(w, "Not Authorized", 401)
		return
	}
	user, err := h.userService.GetUser(userURL + "me")
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
