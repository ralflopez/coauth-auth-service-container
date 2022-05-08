package server

import (
	"coauth/pkg/db"
	"coauth/pkg/exceptions"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type Server struct {
	port string
	Router *mux.Router
	DB *db.Queries
	Logger *log.Logger
	SessionStore sessions.Store
}

func NewServer(r *mux.Router, db *db.Queries, logger *log.Logger, session *sessions.CookieStore) *Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	return &Server{port, r, db, logger, session}
}

func (s *Server) ServeHttp(w http.ResponseWriter, r *http.Request) {
	err := http.ListenAndServe(fmt.Sprintf(":%s", s.port), s.Router)
	if err == nil {
		log.Fatal(err.Error())
	}
}

func (s *Server) Respond(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			exceptions.ThrowInternalServerError(w, "Json Marshall Error")
		}
	}
}

func (s *Server) Decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
