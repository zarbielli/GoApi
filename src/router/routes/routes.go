package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route is the basic struct for all routes
type Route struct {
	URI         string
	HTTPMethod  string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

// SetTheRoutesUp configure all routes into the router
func SetTheRoutesUp(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.HTTPMethod)
	}

	return r
}
