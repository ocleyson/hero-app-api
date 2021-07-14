package tests

import (
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
