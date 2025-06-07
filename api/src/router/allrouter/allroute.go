package allrouter

import (
	"net/http"

	"github.com/gorilla/mux"
)

// AllRoute represents a single route in the API.
type AllRoute struct {
	URI            string
	Method         string
	Function       func(w http.ResponseWriter, r *http.Request)
	Authentication bool
}

// // RouteConfiguration configures routes and return them.
func RouteConfiguration(r *mux.Router) *mux.Router {
	router := userRoutes

	for _, route := range router {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
