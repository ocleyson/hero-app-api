package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ocleyson/hero-app-api/types"
	"github.com/ocleyson/hero-app-api/utils"
)

func SearchHeroByName(res http.ResponseWriter, req *http.Request) {
	name := mux.Vars(req)["name"]

	var searchHeroRes types.SearchHeroesRes

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

	json.NewEncoder(res).Encode(searchHeroRes)
}
