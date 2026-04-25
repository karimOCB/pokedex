package main

import (
	"fmt"
)

func commandExplore(cfg *config, parameters ...string) error {
	if len(parameters) == 0 {
		fmt.Println("A location area to explore is needed. Try again.")
		return nil
	}
	fmt.Printf("Exploring %s...", parameters[0])

	return nil
}
