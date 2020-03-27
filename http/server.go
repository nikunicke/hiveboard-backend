package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/nikunicke/hiveboard/authorize"
)

var baseURL = "https://api.intra.42.fr/v2/"

func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	router := httprouter.New()
	router.GET("/", handleHome)
	router.GET("/login/", handleLogin)
	router.GET("/callback/", handleCallback)
	router.GET("/api/events/", handleEvents)
	router.GET("/api/events/:id", handleEvents)
	router.GET("/api/user/", handleUser)
	log.Println("Server running on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func handleHome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	const html = `
	<body><center>
		<a href="/login/">Login</a>
	</center></body>
	`
	fmt.Fprintf(w, html)
}

func handleLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	url := authorize.GetURL()
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	authorize.GetToken(r.FormValue("code"), r.FormValue("state"))
	fmt.Println("Login successful")
	http.Redirect(w, r, "/api/events/", http.StatusPermanentRedirect)
}

func handleEvents(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	data, err := GetEvents(baseURL + "events/" + p.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func handleUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := GetUser(baseURL + "me")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
