package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("A location area to explore is needed. Try again.")
		return nil
	}
	fmt.Printf("Exploring %s...\n", args[0])

	encounters, err := cfg.Client.LocationExplore(args[0])
	if err != nil {
		return err
	}

	if len(encounters.PokemonEncounters) == 0 {
		fmt.Println("No pokemon found in this area")
		return nil
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range encounters.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)

	}

	return nil
}
