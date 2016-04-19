package web

import (
	"net/http"

	"github.com/lukad/helix/web/assets"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}

func (s *Server) ListenAndServe(address string) error {
	return http.ListenAndServe(address, s.router)
}

func (s *Server) mailsIndex(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`[{"id":0,"from":"from@example.com","to":"to@example.com","subject":"foo","body":"bar"}]`))
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

func NewServer() *Server {
	r := mux.NewRouter()
	server := &Server{r}

	r.HandleFunc("/mails", server.mailsIndex).Methods("GET")
	r.HandleFunc("/mails/{id:[0-9]+}", server.mailsShow).Methods("GET")
	r.HandleFunc("/mails/{id:[0-9]+}", server.mailsDelete).Methods("DEL")
	r.HandleFunc("/helix.js", serveAsset("helix.js", "application/javascript"))
	r.HandleFunc("/helix.js.map", serveAsset("helix.js.map", "application/javascript"))
	r.HandleFunc("/", serveAsset("index.html", "text/html"))

	return server
}
