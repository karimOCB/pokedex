package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("A pokemon name to catch is needed. Try again.")
		return nil
	}

	name := args[0]
	_, ok := cfg.CaughtPokemons[name]
	if ok {
		fmt.Printf("You already catched: %s\n", name)
		return nil
	}

	pokemonInfo, err := cfg.Client.PokemonCatch(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	baseExp := pokemonInfo.BaseExperience

	if rand.Intn(baseExp) < 35 {
		fmt.Printf("%s was caught!\n", name)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.CaughtPokemons[name] = pokemonInfo
		return nil
	}
	fmt.Printf("%s escaped!\n", name)
	return nil
}
