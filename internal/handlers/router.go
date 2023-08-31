package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/neglarken/dynamic_user_segmentation_service/docs"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/service"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Handler struct {
	service *service.Service
	Logger  *logrus.Logger
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
		Logger:  logrus.New(),
	}
}

func NewRouter(h *Handler) *mux.Router {
	router := mux.NewRouter()

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	// router.Handle("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))

	router.Use(h.logRequest)
	router.HandleFunc("/users/", h.CreateUser()).Methods("POST")
	router.HandleFunc("/slugs/", h.CreateSlug()).Methods("POST")
	router.HandleFunc("/slugs/", h.DeleteSlug()).Methods("DELETE")
	router.HandleFunc("/slugsUsers/", h.AddUserInSlugs()).Methods("PUT")
	router.HandleFunc("/slugsUsers/", h.GetUsersSlugs()).Methods("GET")
	router.HandleFunc("/records/", h.GetRecordsByYM()).Methods("GET")
	router.PathPrefix("/files/").Handler(h.files())

	return router
}

func (h *Handler) error(w http.ResponseWriter, r *http.Request, httpCode int, err error) {
	h.respond(w, r, httpCode, map[string]string{"error": err.Error()})
}

func (h *Handler) respond(w http.ResponseWriter, r *http.Request, httpCode int, data interface{}) {
	w.WriteHeader(httpCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
