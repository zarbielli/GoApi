package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// CreateRouter generate the api Router, with
// the routes created
func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.SetTheRoutesUp(r)
}
