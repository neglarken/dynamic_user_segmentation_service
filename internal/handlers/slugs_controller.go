package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/neglarken/dynamic_user_segmentation_service/internal/entity"
)

func (h *Handler) CreateSlug() http.HandlerFunc {
	type request struct {
		Title string `json:"title"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}
		s := &entity.Slugs{
			Title: req.Title,
		}
		if err := h.service.Slugs.Create(req.Title); err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}

		h.respond(w, r, http.StatusOK, s)
	}
}

func (h *Handler) DeleteSlug() http.HandlerFunc {
	type request struct {
		Title string `json:"title"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}
		s := &entity.Slugs{
			Title: req.Title,
		}
		if err := h.service.Slugs.Delete(req.Title); err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}

		h.respond(w, r, http.StatusOK, s)
	}
}
