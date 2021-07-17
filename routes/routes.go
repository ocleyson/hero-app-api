package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ocleyson/hero-app-api/controllers"
	"github.com/rs/cors"
)

func Routes() http.Handler {
	var routes = mux.NewRouter().StrictSlash(true)

	var SearchHeroByName = controllers.SearchHeroByName
	var StoreHero = controllers.StoreHero
	var IndexHeroes = controllers.IndexHeroes
	var IndexGoodHeroes = controllers.IndexGoodHeroes
	var IndexBadHeroes = controllers.IndexBadHeroes
	var ShowHero = controllers.ShowHero
	var DeleteHero = controllers.DeleteHero

	routes.HandleFunc("/search/{name}", SearchHeroByName).Methods("GET")
	routes.HandleFunc("/heroes", IndexHeroes).Methods("GET")
	routes.HandleFunc("/heroes", StoreHero).Methods("POST")
	routes.HandleFunc("/heroes/goods", IndexGoodHeroes).Methods("GET")
	routes.HandleFunc("/heroes/bads", IndexBadHeroes).Methods("GET")
	routes.HandleFunc("/heroes/{id}", ShowHero).Methods("GET")
	routes.HandleFunc("/heroes/{id}", DeleteHero).Methods("DELETE")

	appCors := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE"},
	})

	handler := appCors.Handler(routes)

	return handler
}
