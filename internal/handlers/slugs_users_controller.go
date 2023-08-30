package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/neglarken/dynamic_user_segmentation_service/internal/entity"
)

func (h *Handler) AddUserInSlugs() http.HandlerFunc {
	type request struct {
		TitleAdd    []string `json:"title_add"`
		TitleDelete []string `json:"title_delete"`
		Id          int      `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}

		if len(req.TitleAdd) > 0 {
			for i := 0; i < len(req.TitleAdd); i++ {
				slug, err := h.service.Slugs.GetSlugIdBySlugTitle(req.TitleAdd[i])
				if err != nil {
					h.error(w, r, http.StatusInternalServerError, err)
					return
				}
				if err := h.service.SlugsUsers.Add(&entity.SlugsUsers{
					UserId: req.Id,
					SlugId: slug.Id,
				}); err != nil {
					h.error(w, r, http.StatusInternalServerError, err)
					return
				}
				rec := &entity.Records{
					UserId:    req.Id,
					SlugTitle: req.TitleAdd[i],
					Operation: "create",
				}
				if err := h.service.Records.Create(rec); err != nil {
					h.error(w, r, http.StatusInternalServerError, err)
					return
				}
			}
		}

		if len(req.TitleDelete) > 0 {
			for i := 0; i < len(req.TitleDelete); i++ {
				slug, err := h.service.Slugs.GetSlugIdBySlugTitle(req.TitleDelete[i])
				if err != nil {
					h.error(w, r, http.StatusInternalServerError, err)
					return
				}
				if err := h.service.SlugsUsers.Delete(&entity.SlugsUsers{
					UserId: req.Id,
					SlugId: slug.Id,
				}); err != nil {
					h.error(w, r, http.StatusInternalServerError, err)
					return
				}
				rec := &entity.Records{
					UserId:    req.Id,
					SlugTitle: req.TitleDelete[i],
					Operation: "delete",
				}
				if err := h.service.Records.Create(rec); err != nil {
					h.error(w, r, http.StatusInternalServerError, err)
					return
				}
			}

		}

		type Response struct {
			Status string `json:"status"`
		}

		h.respond(w, r, http.StatusOK, Response{Status: "done :)"})
	}
}

func (h *Handler) GetUsersSlugs() http.HandlerFunc {
	type request struct {
		Id int `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}
		slugsUsers, err := h.service.SlugsUsers.GetSlugIdsByUserId(req.Id)
		if err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}
		slugs := make([]string, 0, len(slugsUsers))
		for i := 0; i < len(slugsUsers); i++ {
			title, err := h.service.Slugs.GetTitleById(slugsUsers[i].SlugId)
			if err != nil {
				h.error(w, r, http.StatusInternalServerError, err)
				return
			}
			slugs = append(slugs, title.Title)
		}

		type Response struct {
			UserId int      `json:"user_id"`
			Slugs  []string `json:"slugs"`
		}

		h.respond(w, r, http.StatusOK, Response{UserId: req.Id, Slugs: slugs})
	}
}
