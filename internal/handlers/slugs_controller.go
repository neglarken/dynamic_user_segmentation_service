package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/neglarken/dynamic_user_segmentation_service/internal/entity"
)

func (h *Handler) CreateSlug() http.HandlerFunc {
	type request struct {
		Title string `json:"title"`
		Part  int    `json:"part"`
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
		if err := h.service.Slugs.Create(s); err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}
		if req.Part != 0 {
			countUsers, err := h.service.Users.GetCount()
			if err != nil {
				h.error(w, r, http.StatusInternalServerError, err)
				return
			}
			count := countUsers * req.Part / 100
			users, err := h.service.Users.GetNumOfRandom(count)
			if err != nil {
				h.error(w, r, http.StatusInternalServerError, err)
				return
			}
			for i := 0; i < len(users); i++ {
				if err := h.service.SlugsUsers.Add(&entity.SlugsUsers{
					UserId: users[i].Id,
					SlugId: s.Id,
				}); err != nil {
					h.error(w, r, http.StatusInternalServerError, err)
					return
				}
				rec := &entity.Records{
					UserId:    users[i].Id,
					SlugTitle: s.Title,
					Operation: "create",
				}
				if err := h.service.Records.Create(rec); err != nil {
					h.error(w, r, http.StatusInternalServerError, err)
					return
				}
			}
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
		if err := h.service.Slugs.Delete(s); err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}

		type response struct {
			Status string `json:"status"`
		}

		h.respond(w, r, http.StatusOK, response{Status: "done"})
	}
}
