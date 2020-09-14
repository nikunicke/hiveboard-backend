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
	"github.com/rs/cors"
)

var baseURL = "https://api.intra.42.fr/v2/"

// Server includes the services and some settings
type Server struct {
	ln net.Listener

	// EventService hiveboard.EventService
	ES hiveboard.EventService2

	UserService hiveboard.UserService

	Addr        string
	Host        string
	Autocert    bool
	Recoverable bool
	LogOutput   io.Writer
}

// NewServer defines a new server with the Recoverable variable set as true
func NewServer() *Server {
	return &Server{Recoverable: true}
}

// Open opens a connection to the server
func (s *Server) Open() error {
	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	s.ln = ln
	go http.Serve(s.ln, s.router())
	return nil
}

// Close takes down the server
func (s *Server) Close() error {
	if s.ln != nil {
		s.ln.Close()
	}
	return nil
}

// URL ...
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
	r.Use(cors.Default().Handler)
	r.Route("/", func(r chi.Router) {
		r.Get("/", handleIndex)
		r.Get("/login/", handleLogin)
		r.Get("/callback/", handleCallback)
		r.Mount("/api/events", s.eventHandler())
		r.Mount("/api/user", s.userHandler())
	})
	return r
}

func (s *Server) eventHandler() *eventHandler {
	h := newEventHandler()
	h.baseURL = s.URL()
	h.eventService2 = s.ES
	return h
}

func (s *Server) userHandler() *userHandler {
	h := newUserHandler()
	h.baseURL = s.URL()
	h.userService = s.UserService
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
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	authorize.GetToken(r.FormValue("code"), r.FormValue("state"))
	// fmt.Println(hiveboard.OauthToken)
	http.Redirect(w, r, "/api/events/", http.StatusTemporaryRedirect)
}
