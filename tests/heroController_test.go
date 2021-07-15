package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ocleyson/hero-app-api/routes"
	"github.com/ocleyson/hero-app-api/services"
	"github.com/stretchr/testify/assert"
)

func clearTable() {
	services.DB.Exec("DELETE FROM heros")
}

func setup() {
	services.ConnectDatabase()
	clearTable()
	http.Handle("/", routes.Routes())
}

func shutdown() {
	http.DefaultServeMux = new(http.ServeMux)
}

func createHero() (*httptest.ResponseRecorder, error) {
	payload := []byte(`{"id": "33", "name": "batman", "fullName": "bruce wayne", "intelligence": "22", "power": "32", "occupation": "businesman", "imageUrl": "http://url", "groupAffiliation": "batman family", "relatives": "bruce wayne (biological father)", "alignment": "good"}`)

	req, err := http.NewRequest("POST", "/heroes", bytes.NewBuffer(payload))

	res := httptest.NewRecorder()

	routes.Routes().ServeHTTP(res, req)

	return res, err
}

func TestSearchHeroByName(t *testing.T) {
	setup()

	req, err := http.NewRequest("GET", "/search/batman", nil)

	res := httptest.NewRecorder()

	if err != nil {
		t.Fatalf(`test search hero error: %q`, err)
	}

	routes.Routes().ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "it should search hero by its name")

	shutdown()
}

func TestStoreHero(t *testing.T) {
	setup()

	res, err := createHero()

	if err != nil {
		t.Fatalf(`create hero error: %q`, err)
	}

	assert.Equal(t, http.StatusOK, res.Code, "it should create a hero")

	shutdown()
}

func TestIndexHeroes(t *testing.T) {
	setup()

	req, err := http.NewRequest("GET", "/heroes", nil)

	res := httptest.NewRecorder()

	if err != nil {
		t.Fatalf(`get all stored heroes: %q`, err)
	}

	routes.Routes().ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "it should get all stored heroes")

	shutdown()
}

func TestIndexGoodHeroes(t *testing.T) {
	setup()

	req, err := http.NewRequest("GET", "/heroes/goods", nil)

	res := httptest.NewRecorder()

	if err != nil {
		t.Fatalf(`get all stored good heroes: %q`, err)
	}

	routes.Routes().ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "it should get all stored good heroes")

	shutdown()
}

func TestIndexBadHeroes(t *testing.T) {
	setup()

	req, err := http.NewRequest("GET", "/heroes/bads", nil)

	res := httptest.NewRecorder()

	if err != nil {
		t.Fatalf(`get all stored bad heroes: %q`, err)
	}

	routes.Routes().ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "it should get all stored bad heroes")

	shutdown()
}

func TestShowHero(t *testing.T) {
	setup()

	_, errHero := createHero()

	if errHero != nil {
		t.Fatalf(`failed to create hero: %q`, errHero)
	}

	req, err := http.NewRequest("GET", "/heroes/33", nil)

	res := httptest.NewRecorder()

	if err != nil {
		t.Fatalf(`get hero by id: %q`, err)
	}

	routes.Routes().ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "it should get a hero by its id")

	shutdown()
}
