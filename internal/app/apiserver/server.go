package apiserver

import (
	"github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/kantegory/http-rest-api/internal/app/store"
	"net/http"
)

// server ...
type server struct {
	router *mux.Router
	logger *logrus.Logger
	store store.Store
}

// newServer() ...
func newServer(store store.Store) *server {
	s := &server {
		router: mux.NewRouter(),
		logger: logrus.New(),
		store: store,
	}

	s.configureRouter()

	return s
}

// ServeHTTP() ...
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// configureRouter() ...
func (s * server) configureRouter() {
	// ...
}
