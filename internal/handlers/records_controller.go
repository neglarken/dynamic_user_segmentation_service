package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

// @Summary Get record by year-month
// @Description Get link to records by year-month
// @Tags record
// @Accept json
// @Produce json
// @Param input body handlers.GetRecordsByYM.Request true "date year-month"
// @Success 301
// @Failure 500
// @Router /records/ [get]
func (h *Handler) GetRecordsByYM() http.HandlerFunc {
	type Request struct {
		Date string `json:"date"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &Request{}
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

// @Summary Get link to records
// @Description Get link to records
// @Tags record
// @Success 301 {integer} integer
// @Failure 404
// @Router /files/ [get]
func (h Handler) files() http.Handler {
	return http.StripPrefix("/files", http.FileServer(http.Dir("./records")))
}
