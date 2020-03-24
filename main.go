package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

type Event struct {
	Name string
}

var token *oauth2.Token

func handleHome(w http.ResponseWriter, r *http.Request) {
	var html = `
	<body><center>
		<a href="/login/">Login</a>
	</center></body>
	`
	fmt.Fprintf(w, html)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := getAuthorization()
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleOauthCallback(w http.ResponseWriter, r *http.Request) {
	var err error

	token, err = getAccessToken(r.FormValue("code"), r.FormValue("state"))
	if err != nil {
		// http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		http.Error(w, err.Error(), 500)
	}
	http.Redirect(w, r, "/api/events/", http.StatusPermanentRedirect)
}

func handleEvents(w http.ResponseWriter, r *http.Request) {
	url := "https://api.intra.42.fr/v2/events"
	var err error
	if token == nil {
		err := errors.New("Authorization token missing")
		http.Error(w, err.Error(), 401)
		return
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", string([]byte(body)))
	// var events []Event
	// err = json.Unmarshal(body, &events)
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }
	// fmt.Printf("name: %s\n", events[0].Name)
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(data)
}

func main() {
	port := "3000"

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login/", handleLogin)
	http.HandleFunc("/callback/", handleOauthCallback)
	http.HandleFunc("/api/events/", handleEvents)
	fmt.Println("Server up and running on port 3000")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
