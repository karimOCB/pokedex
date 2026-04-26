package main

import (
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
		fmt.Printf("Couln't create new readline %s", err)
		return
	}

	cfg := config{
		Next:           nil,
		Previous:       nil,
		Client:         client,
		CaughtPokemons: map[string]pokeapi.Pokemon{},
		RlInstance:     rl,
	}
	defer cfg.RlInstance.Close()

	startRepl(&cfg)
}
