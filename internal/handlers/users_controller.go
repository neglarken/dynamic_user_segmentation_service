package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/neglarken/dynamic_user_segmentation_service/internal/entity"
)

// @Summary Create user
// @Description Create user
// @Tags Users
// @Accept json
// @Produce json
// @Param input body handlers.CreateUser.Request true "input id"
// @Success 200 {object} entity.Users
// @Failure 500
// @Router /users/ [post]
func (h *Handler) CreateUser() http.HandlerFunc {
	type Request struct {
		Id int `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}
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
