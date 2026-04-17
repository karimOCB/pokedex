package main

import (
	"net/http"
)

type Client struct {
	httpClient http.Client
}

func commandMap(cfg *config) error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://pokeapi.co/api/v2/location-area/{id or name}/", nil)

	resp, err := client.Do(req)
}
