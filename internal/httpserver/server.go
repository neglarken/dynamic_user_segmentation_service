package httpserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/handlers"
)

type Server struct {
	router *mux.Router
}

func NewServer(handler handlers.Handler) *Server {
	s := &Server{
		router: handlers.NewRouter(&handler),
	}

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) NewSubRouter(str string) *mux.Router {
	return s.router.PathPrefix(str).Subrouter()
}
