package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgconn"
	"github.com/ocleyson/hero-app-api/models"
	"github.com/ocleyson/hero-app-api/services"
	"github.com/ocleyson/hero-app-api/types"
	"github.com/ocleyson/hero-app-api/utils"
)

func SearchHeroByName(res http.ResponseWriter, req *http.Request) {
	name := mux.Vars(req)["name"]

	var searchHeroRes types.SearchHeroesRes

	var heroes []models.Hero

	var hero models.Hero

	SUPER_HERO_API_KEY := utils.GetEnvVar("SUPER_HERO_API_KEY")

	url := "https://superheroapi.com/api/" + SUPER_HERO_API_KEY + "/search/" + name

	response, errRes := http.Get(url)

	if errRes != nil {
		http.Error(res, errRes.Error(), http.StatusBadRequest)
		return
	}

	reqBody, errBody := ioutil.ReadAll(response.Body)

	if errBody != nil {
		http.Error(res, errBody.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &searchHeroRes)

	for i := 0; i < len(searchHeroRes.Results); i++ {
		hero.Id = searchHeroRes.Results[i].Id
		hero.Name = searchHeroRes.Results[i].Name
		hero.FullName = searchHeroRes.Results[i].Biography.FullName
		hero.Intelligence = searchHeroRes.Results[i].Powerstats.Intelligence
		hero.Power = searchHeroRes.Results[i].Powerstats.Power
		hero.Occupation = searchHeroRes.Results[i].Work.Occupation
		hero.ImageUrl = searchHeroRes.Results[i].Image.Url
		hero.GroupAffiliation = searchHeroRes.Results[i].Connections.GroupAffiliation
		hero.Relatives = searchHeroRes.Results[i].Connections.Relatives
		hero.Alignment = searchHeroRes.Results[i].Biography.Alignment

		heroes = append(heroes, hero)
	}

	json.NewEncoder(res).Encode(heroes)
}

func StoreHero(res http.ResponseWriter, req *http.Request) {

	reqBody, _ := ioutil.ReadAll(req.Body)

	var hero models.Hero

	json.Unmarshal(reqBody, &hero)

	result := services.DB.Create(&hero)

	var perr *pgconn.PgError

	if result.Error != nil {
		errors.As(result.Error, &perr)

		if perr.Code != "23505" {
			http.Error(res, result.Error.Error(), http.StatusBadRequest)
			return
		}
	}

	json.NewEncoder(res).Encode(hero)
}

func IndexHeroes(res http.ResponseWriter, req *http.Request) {
	var heroes []models.Hero

	result := services.DB.Find(&heroes)

	if result.Error != nil {
		http.Error(res, result.Error.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(res).Encode(heroes)
}

func IndexGoodHeroes(res http.ResponseWriter, req *http.Request) {
	var heroes []models.Hero

	result := services.DB.Where("alignment = ?", "good").Find(&heroes)

	if result.Error != nil {
		http.Error(res, result.Error.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(res).Encode(heroes)
}

func IndexBadHeroes(res http.ResponseWriter, req *http.Request) {
	var heroes []models.Hero

	result := services.DB.Where("alignment = ?", "bad").Find(&heroes)

	if result.Error != nil {
		http.Error(res, result.Error.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(res).Encode(heroes)
}

func ShowHero(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	var getHeroRes types.GetHeroRes

	var hero models.Hero

	SUPER_HERO_API_KEY := utils.GetEnvVar("SUPER_HERO_API_KEY")

	url := "https://superheroapi.com/api/" + SUPER_HERO_API_KEY + "/" + id

	response, errRes := http.Get(url)

	if errRes != nil {
		http.Error(res, errRes.Error(), http.StatusBadRequest)
		return
	}

	reqBody, errBody := ioutil.ReadAll(response.Body)

	if errBody != nil {
		http.Error(res, errBody.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &getHeroRes)

	hero.Id = getHeroRes.Id
	hero.Name = getHeroRes.Name
	hero.FullName = getHeroRes.Biography.FullName
	hero.Intelligence = getHeroRes.Powerstats.Intelligence
	hero.Power = getHeroRes.Powerstats.Power
	hero.Occupation = getHeroRes.Work.Occupation
	hero.ImageUrl = getHeroRes.Image.Url
	hero.GroupAffiliation = getHeroRes.Connections.GroupAffiliation
	hero.Relatives = getHeroRes.Connections.Relatives
	hero.Alignment = getHeroRes.Biography.Alignment

	json.NewEncoder(res).Encode(hero)
}

func DeleteHero(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	var hero models.Hero

	result := services.DB.Delete(&hero, id)

	if result.Error != nil {
		http.Error(res, result.Error.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(res).Encode(id)
}
