package web

import (
	"encoding/json"
	"net/http"

	"github.com/lukad/helix/store"
	"github.com/lukad/helix/web/assets"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	store  store.Store
}

type listResponse struct {
	Items []store.Mail `json:"items"`
}

type subject struct {
	Id      int    `json:"id"`
	Subject string `json:"subject"`
}

type subjectsResponse struct {
	Items []subject
}

func (s *Server) ListenAndServe(address string) error {
	return http.ListenAndServe(address, s.router)
}

func (s *Server) mailsIndex(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response listResponse
	for _, item := range s.store.All() {
		response.Items = append(response.Items, item)
	}
	b, _ := json.Marshal(response)
	w.Write(b)
}

func (s *Server) mailsSubjects(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response subjectsResponse
	for _, item := range s.store.All() {
		sub := subject{
			Id:      item.Id,
			Subject: item.Header["Subject"][0],
		}
		response.Items = append(response.Items, sub)
	}
	b, _ := json.Marshal(response)
	w.Write(b)
}

func (s *Server) mailsShow(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"id":0,"from":"from@example.com","to":"to@example.com","subject":"foo","body":"bar"}`))
}

func (s *Server) mailsDelete(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
}

func serveAsset(name string, contentType string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if data, err := assets.Asset(name); err == nil {
			w.Header().Set("Content-Type", contentType)
			w.Write(data)
		} else {
			http.Error(w, "not found", http.StatusNotFound)
		}
	}
}

func NewServer(s store.Store) *Server {
	r := mux.NewRouter()
	server := &Server{
		router: r,
		store:  s,
	}

	r.HandleFunc("/mails", server.mailsIndex).Methods("GET")
	r.HandleFunc("/mails/subjects", server.mailsSubjects).Methods("GET")
	r.HandleFunc("/mails/{id:[0-9]+}", server.mailsShow).Methods("GET")
	r.HandleFunc("/mails/{id:[0-9]+}", server.mailsDelete).Methods("DEL")
	r.HandleFunc("/helix.js", serveAsset("helix.js", "application/javascript"))
	r.HandleFunc("/helix.js.map", serveAsset("helix.js.map", "application/javascript"))
	r.HandleFunc("/helix.css", serveAsset("helix.css", "text/css"))
	r.HandleFunc("/helix.css.map", serveAsset("helix.css.map", "text/css"))
	r.HandleFunc("/", serveAsset("index.html", "text/html"))

	return server
}
