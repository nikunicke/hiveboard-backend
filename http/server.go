package http

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/nikunicke/hiveboard"
	"github.com/nikunicke/hiveboard/authorize"
)

var baseURL = "https://api.intra.42.fr/v2/"

type Server struct {
	ln net.Listener

	EventService hiveboard.EventService

	Addr        string
	Host        string
	Autocert    bool
	Recoverable bool
	LogOutput   io.Writer
}

func NewServer() *Server {
	return &Server{Recoverable: true}
}

func (s *Server) Open() error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	s.ln = ln
	fmt.Println("Server running on: " + s.Addr)
	http.Serve(s.ln, s.router())
	return nil
}

func (s *Server) Close() error {
	if s.ln != nil {
		s.ln.Close()
	}
	return nil
}

func (s *Server) URL() url.URL {
	if s.ln == nil {
		return url.URL{}
	}
	return url.URL{
		Scheme: "http",
		Host:   s.ln.Addr().String(),
	}
}

func (s *Server) router() http.Handler {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Get("/", handleIndex)
		r.Get("/login/", handleLogin)
		r.Get("/callback/", handleCallback)
		r.Mount("/api/events/", s.eventHandler())
	})
	return r
}

func (s *Server) eventHandler() *eventHandler {
	h := newEventHandler()
	h.baseURL = s.URL()
	h.eventService = s.EventService
	return h
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	const html = `
	<body><center>
		<a href="/login/">Login</a>
	</center></body>
	`
	fmt.Fprintf(w, html)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := authorize.GetURL()
	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	authorize.GetToken(r.FormValue("code"), r.FormValue("state"))
	http.Redirect(w, r, "/api/events/", http.StatusPermanentRedirect)
}

// func Run() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "3000"
// 	}
// 	router := httprouter.New()
// 	router.GET("/", handleHome)
// 	router.GET("/login/", handleLogin)
// 	router.GET("/callback/", handleCallback)
// 	router.GET("/api/events/", handleEvents)
// 	router.GET("/api/events/:id", handleEvents)
// 	router.GET("/api/user/", handleUser)
// 	log.Println("Server running on port: " + port)
// 	log.Fatal(http.ListenAndServe(":"+port, router))
// }

// func handleHome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	const html = `
// 	<body><center>
// 		<a href="/login/">Login</a>
// 	</center></body>
// 	`
// 	fmt.Fprintf(w, html)
// }

// func handleLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	url := authorize.GetURL()
// 	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
// }

// func handleCallback(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	authorize.GetToken(r.FormValue("code"), r.FormValue("state"))
// 	fmt.Println("Login successful")
// 	http.Redirect(w, r, "/api/events/", http.StatusPermanentRedirect)
// }

// func handleEvents(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 	data, err := GetEvents(baseURL + "events/" + p.ByName("id"))
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(data)
// }

// func handleUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	data, err := GetUser(baseURL + "me")
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(data)
// }
