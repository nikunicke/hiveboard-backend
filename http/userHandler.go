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
	// h.router.Get("/sub", h.eventsUsers)
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

// func (h *userHandler) eventsUsers(w http.ResponseWriter, r *http.Request) {
// 	// userID := chi.URLParam(r, "userID") + "/"
// 	fmt.Println(hiveboard.OauthToken.AccessToken)
// 	res, err := hiveboard.Client.PostForm("https://profile.intra.42.fr/events/4376/events_users", url.Values{"_method": {"post"}, "authenticity_token": {"0q06d1mkw%2F7r4KsIEu3Cj2yG%2BjvwmDf8VAom2xBXWmqIJ%2BwEmT069mQCc0iXdJqvAriN3LYWqbQF6GEE0NZUjw%3D%3D"}})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Fprintf(w, string(body))
// }
