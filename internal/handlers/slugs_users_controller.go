package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) AddUserInSlugs() http.HandlerFunc {
	type request struct {
		Title []string `json:"title"`
		Id    int      `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}
		if err := h.service.SlugsUsers.Add(req.Title, req.Id); err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}

		type Response struct {
			Status string `json:"status"`
		}

		h.respond(w, r, http.StatusOK, Response{Status: "done"})
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
		slugs, err := h.service.SlugsUsers.Get(req.Id)
		if err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}

		h.respond(w, r, http.StatusOK, slugs)
	}
}
