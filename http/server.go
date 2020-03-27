package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nikunicke/hiveboard/authorize"
)

var baseURL = "https://api.intra.42.fr/v2/"

func Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login/", handleLogin)
	http.HandleFunc("/callback/", handleCallback)
	http.HandleFunc("/api/events/", handleEvents)
	http.HandleFunc("/api/user/", handleUser)
	log.Println("Server running on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	const html = `
	<body><center>
		<a href="/login/">Login</a>
	</center></body>
	`
	fmt.Fprintf(w, html)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := authorize.GetURL()
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	authorize.GetToken(r.FormValue("code"), r.FormValue("state"))
	fmt.Println("Login successful")
	http.Redirect(w, r, "/api/events/", http.StatusPermanentRedirect)
}

func handleEvents(w http.ResponseWriter, r *http.Request) {
	data, err := GetEvents(baseURL + "events")
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	data, err := GetUser(baseURL + "me")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
