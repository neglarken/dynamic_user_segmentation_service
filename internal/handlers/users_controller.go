package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/neglarken/dynamic_user_segmentation_service/internal/entity"
)

func (h *Handler) CreateUser() http.HandlerFunc {
	type request struct {
		Id int `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &entity.Users{
			Id: req.Id,
		}
		if err := h.service.Users.Create(u); err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}

		h.respond(w, r, http.StatusOK, u)
	}
}
