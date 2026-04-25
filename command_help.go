package main

import "fmt"

func commandHelp(cfg *config, parameters ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	registry := getCommand()

	for key, value := range registry {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}
