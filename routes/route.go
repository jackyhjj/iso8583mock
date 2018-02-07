// Package routes is the route & midlleware service package
package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// Route http request pattern
type Route struct {
	HealthCheckController interfaces.IHealthCheckController
}

// NewRoute instances
func NewRoute() *Route {
	return &Route{}
}

// GetRouter Build the all router
func (r *Route) GetRouter() http.Handler {

	router := mux.NewRouter().StrictSlash(true)

	rootRoute := router.PathPrefix("/v1/").Subrouter()
	//ptpos := rootRoute.PathPrefix("/ptpos/").Subrouter()

	rootRoute.HandleFunc("/healthcheck", r.HealthCheckController.HeartBeat).Methods("GET")

	n := negroni.New()
	n.Use(NewLoggerMiddleware())
	n.UseHandler(rootRoute)

	return n
}
