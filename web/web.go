package web

import (
	"net/http"

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

func NewServer() *Server {
	r := mux.NewRouter()
	server := &Server{r}

	r.HandleFunc("/mails", server.mailsIndex).Methods("GET")
	r.HandleFunc("/mails/{id:[0-9]+}", server.mailsShow).Methods("GET")
	r.HandleFunc("/mails/{id:[0-9]+}", server.mailsDelete).Methods("DEL")

	return server
}
