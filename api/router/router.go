package router

import (
	"github.com/gorilla/mux"
	"log"
	"mpg-tracker/api/repositories/fill_up_repository"
	"mpg-tracker/api/utilities"
	"net/http"
)

type Config struct {
	ConnectionString string
}

func configureRoutes(name string, router *mux.Router, routes Routes) {
	log.Printf("Configuring routes for %s", name)

	for _, route := range routes {
		log.Printf("%-6s %-16s %s\n", route.Method, route.Name, route.Pattern)

		var handler http.Handler
		handler = route.HandlerFunc
		handler = utilities.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	log.Println()
}

func NewRouter(config Config) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Health
	configureRoutes("Health", router, MakeHealthRoutes())

	// FillUps
	repo := fill_up_repository.New(config.ConnectionString)
	configureRoutes("Fill Ups", router, MakeFillUpRoutes(repo))

	return router
}
