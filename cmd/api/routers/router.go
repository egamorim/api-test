package routers

import (
	"github.com/egamorim/api-test/cmd/api/handlers"
	"github.com/gorilla/mux"
)

func Router(h *handlers.Handler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health", h.HandleHealthCheck).Methods("GET")

	return router
}
