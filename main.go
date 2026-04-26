package main

import (
	"math/rand/v2"
	"fmt"
	"time"

	"github.com/chzyer/readline"
	"github.com/karimOCB/pokedex/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)

	rl, err := readline.NewEx(&readline.Config{
		Prompt: "Pokedex > ",
	})
	if err != nil {
		fmt.Printf("Couldn't create new readline %s", err)
		return
	}

	cfg := config{
		Next:           nil,
		Previous:       nil,
		Client:         client,
		CaughtPokemons: map[string]pokeapi.Pokemon{},
		RlInstance:     rl,
		CurrentOpponent: nil,
	}
	defer cfg.RlInstance.Close()
	
	go randomOpponent(&cfg)
	startRepl(&cfg)
}

func randomOpponent(cfg *config) {
	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		if len(cfg.CaughtPokemons) == 0 || cfg.CurrentOpponent != nil {
			continue
		} 

		pokemonInfo, err := cfg.Client.PokemonCatch(rand.IntN(1025))
		if err != nil {
			continue
		}

		cfg.Mux.Lock()
		cfg.CurrentOpponent = &pokemonInfo
		fmt.Fprintln(cfg.RlInstance, "\n A wild Pokemon has appeared!")
		cfg.Mux.Unlock()
	}
}