package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("A pokemon name is needed to inspect. Try again")
	}

	name := args[0]

	val, ok := cfg.CaughtPokemons[name]
	if !ok {
		return fmt.Errorf("You have not caugth %s yet. Try catch command.", name)
	}

	fmt.Printf("Name: %s\n", val.Name)
	fmt.Printf("Height: %d\n", val.Height)
	fmt.Printf("Weight: %d\n", val.Weight)

	fmt.Println("Stats:")
	for _, stat := range val.Stats {
		fmt.Printf("-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, pokemon_type := range val.Types {
		fmt.Printf("- %s\n", pokemon_type.Type.Name)
	}

	return nil
}
