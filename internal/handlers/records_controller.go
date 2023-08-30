package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

func (h *Handler) GetRecordsByYM() http.HandlerFunc {
	type request struct {
		Date string `json:"date"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			h.error(w, r, http.StatusBadRequest, err)
			return
		}
		time, err := time.Parse("2006-01", req.Date)
		if err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}
		err = h.service.Records.GetToCsv(time)
		if err != nil {
			h.error(w, r, http.StatusInternalServerError, err)
			return
		}
		http.Redirect(w, r, "/files/", http.StatusMovedPermanently)
	}
}

func (h Handler) files() http.Handler {
	return http.StripPrefix("/files", http.FileServer(http.Dir("./records")))
}
