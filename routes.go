package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ocleyson/hero-app-api/controllers"
)

func Routes() http.Handler {
	var routes = mux.NewRouter().StrictSlash(true)
	var SearchHeroByName = controllers.SearchHeroByName

	routes.HandleFunc("/search/{name}", SearchHeroByName).Methods("GET")

	return routes
}
