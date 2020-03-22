package main

import (
	"fmt"
	"log"
	"net/http"
)

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
	token, err := getAccessToken(r.FormValue("code"), r.FormValue("state"))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}
	fmt.Fprintf(w, "Sign in successful")
	fmt.Printf("%v\n", token)
}

func main() {
	port := "3000"

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login/", handleLogin)
	http.HandleFunc("/callback/", handleOauthCallback)
	fmt.Println("Server up and running on port 3000")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
