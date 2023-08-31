package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/neglarken/dynamic_user_segmentation_service/internal/entity"
)

// @Summary Create slug
// @Description Create slug
// @Tags slugs
// @Accept json
// @Produce json
// @Param input body handlers.CreateSlug.Request true "input title and part"
// @Success 200 {object} entity.Slugs
// @Failure 500
// @Router /slugs/ [post]
func (h *Handler) CreateSlug() http.HandlerFunc {
	type Request struct {
		Title string `json:"title"`
		Part  int    `json:"part"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}
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

// @Summary Delete slug
// @Description Delete slug
// @Tags slugs
// @Accept json
// @Produce json
// @Param input body handlers.DeleteSlug.Request true "input title"
// @Success 200 {object} handlers.DeleteSlug.Response
// @Failure 500
// @Router /slugs/ [post]
func (h *Handler) DeleteSlug() http.HandlerFunc {
	type Request struct {
		Title string `json:"title"`
	}
	type Response struct {
		Status string `json:"status"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}
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

		h.respond(w, r, http.StatusOK, Response{Status: "done"})
	}
}
