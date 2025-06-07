package router

import (
	"SraReaper/enube-me-api/api/src/router/allrouter"

	"github.com/gorilla/mux"
)

// Route initializes a new router.
func Route() *mux.Router {
	r := mux.NewRouter()
	return allrouter.RouteConfiguration(r)
}
