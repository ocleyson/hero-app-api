package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ocleyson/hero-app-api/controllers"
)

func Routes() http.Handler {
	var routes = mux.NewRouter().StrictSlash(true)
	var SearchHeroByName = controllers.SearchHeroByName
	var StoreHero = controllers.StoreHero
	var IndexHeroes = controllers.IndexHeroes

	routes.HandleFunc("/search/{name}", SearchHeroByName).Methods("GET")
	routes.HandleFunc("/heroes", StoreHero).Methods("POST")
	routes.HandleFunc("/heroes", IndexHeroes).Methods("GET")

	return routes
}
