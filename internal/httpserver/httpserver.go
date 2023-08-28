package httpserver

import (
	"database/sql"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/neglarken/dynamic_user_segmentation_service/config"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/handlers"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/repository"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/service"
	"github.com/neglarken/dynamic_user_segmentation_service/internal/store"
)

func Start(config *config.Config) error {
	db, err := NewDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := store.NewStore(db)
	repos := repository.NewRepository(store)
	serv := service.NewService(*repos)
	handler := handlers.NewHandler(serv)
	s := NewServer(*handler)
	handler.Logger.Infof("Starting server on %s", config.Addr)
	return http.ListenAndServe(config.Addr, s)
}

func NewDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
