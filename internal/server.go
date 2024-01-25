package internal

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	httpd *http.Server
}

func New() *Server {
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
	})

	return &Server{
		httpd: &http.Server{
			Addr:    "0.0.0.0:8080",
			Handler: router,
		},
	}
}

func (server *Server) Start() error {
	return server.httpd.ListenAndServe()
}
