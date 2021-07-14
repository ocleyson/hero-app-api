package main

import (
	"net/http"
)

func main() {
	err := http.ListenAndServe(":3000", Routes())

	if err != nil {
		panic("Failed to serve api.")
	}
}
