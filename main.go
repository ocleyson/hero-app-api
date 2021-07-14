package main

import (
	"net/http"

	"github.com/ocleyson/hero-app-api/services"
)

func main() {
	services.ConnectDatabase()

	err := http.ListenAndServe(":3000", Routes())

	if err != nil {
		panic("Failed to serve api.")
	}
}
