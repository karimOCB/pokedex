package main

import (
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	locations, err := cfg.Client.ListLocations(cfg.Next)

	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, err := cfg.Client.ListLocations(cfg.Previous)

	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous
	return nil
}
