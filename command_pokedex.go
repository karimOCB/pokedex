package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.CaughtPokemons) == 0 {
		fmt.Println("You don't have pokemons yet. Go and catch them.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for key := range cfg.CaughtPokemons {
		fmt.Printf("- %s\n", key)
	}

	return nil
}
