package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/neglarken/dynamic_user_segmentation_service/internal/entity"
)

// @Summary Add users in slugs
// @Description Add users in slugs
// @Tags Segments
// @Accept json
// @Produce json
// @Param input body handlers.AddUserInSlugs.Request true "input [title_add], [title_delete], id, ttl"
// @Success 200 {object} handlers.AddUserInSlugs.Response
// @Failure 500
// @Router /slugsUsers/ [post]
func (h *Handler) AddUserInSlugs() http.HandlerFunc {
	type Request struct {
		TitleAdd    []string `json:"title_add"`
		TitleDelete []string `json:"title_delete"`
		Id          int      `json:"id"`
		Ttl         int      `json:"ttl"`
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
			if req.Ttl > 0 {
				for i := 0; i < len(req.TitleAdd); i++ {
					slug, err := h.service.Slugs.GetSlugIdBySlugTitle(req.TitleAdd[i])
					if err != nil {
						h.error(w, r, http.StatusInternalServerError, err)
						return
					}
					timer := time.NewTimer(time.Duration(req.Ttl) * time.Second)
					go func() {
						<-timer.C
						h.service.SlugsUsers.Delete(&entity.SlugsUsers{
							UserId: req.Id,
							SlugId: slug.Id,
						})
						rec := &entity.Records{
							UserId:    req.Id,
							SlugTitle: slug.Title,
							Operation: "delete",
						}
						if err := h.service.Records.Create(rec); err != nil {
							h.error(w, r, http.StatusInternalServerError, err)
							return
						}
						h.Logger.Infoln("deleted")
					}()
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
		h.respond(w, r, http.StatusOK, Response{Status: "done :)"})
	}
}

// @Summary Get users slugs
// @Description Get users slugs
// @Tags Segments
// @Accept json
// @Produce json
// @Param input body handlers.AddUserInSlugs.Request true "input id"
// @Success 200 {object} handlers.GetUsersSlugs.Response
// @Failure 500
// @Router /slugsUsers/ [get]
func (h *Handler) GetUsersSlugs() http.HandlerFunc {
	type Request struct {
		Id int `json:"id"`
	}
	type Response struct {
		UserId int      `json:"user_id"`
		Slugs  []string `json:"slugs"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}
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

		h.respond(w, r, http.StatusOK, Response{UserId: req.Id, Slugs: slugs})
	}
}
