package handler

import (
	"net/http"
)

func (h *Handler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	err := h.Json(w, http.StatusOK, "API is running")
	if err != nil {
		h.Log("error", err)
	}
}
