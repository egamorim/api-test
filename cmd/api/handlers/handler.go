package handlers

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
}

func (h *Handler) ResponseWithError(w http.ResponseWriter, code int, message string) {
	h.RespondWithJSON(w, code, map[string]string{"error": message})
}

func (h *Handler) RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
