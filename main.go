package main

import (
	"time"

	"github.com/karimOCB/pokedex/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cfg := config{
		Next:     nil,
		Previous: nil,
		Client:   client,
	}

	startRepl(&cfg)
}
