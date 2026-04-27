package main

import (
	"fmt"
)

type battleStats struct {
	attack int
	hp     int
}

func commandBattle(cfg *config, args ...string) error {

	if len(args) == 0 {
		fmt.Println("You need to choose a pokemon from your pokedex to fight.")
		return nil
	}

	cfg.Mux.Lock()
	if cfg.CurrentOpponent == nil {
		cfg.Mux.Unlock()
		fmt.Println("There is no wild Pokemon to battle right now!")
		return nil
	}

	chosenPokemon, ok := cfg.CaughtPokemons[args[0]]

	if !ok {
		fmt.Printf("You don't have %s in your pokedex. Choose one you have.", args[0])
		cfg.Mux.Unlock()
		return nil
	}

	rivalPokemon := cfg.CurrentOpponent

	var chosenBattleStats battleStats
	var rivalBattleStats battleStats

	for _, stat := range chosenPokemon.Stats {
		if stat.Stat.Name == "attack" {
			chosenBattleStats.attack = stat.BaseStat
		}
		if stat.Stat.Name == "hp" {
			chosenBattleStats.hp = stat.BaseStat
		}
	}

	for _, stat := range rivalPokemon.Stats {
		if stat.Stat.Name == "attack" {
			rivalBattleStats.attack = stat.BaseStat
		}
		if stat.Stat.Name == "hp" {
			rivalBattleStats.hp = stat.BaseStat
		}
	}
	cfg.Mux.Unlock()

	for i := 1; ; i++ {
		fmt.Printf("\nRound #%d\n\n", i)

		if i == 1 {
			fmt.Println("The enemy attacks first by surprise.")
		}
		fmt.Printf("Enemy attack #%d\n", i)

		chosenBattleStats.hp -= rivalBattleStats.attack
		fmt.Printf("Now your HP is %d\n\n", chosenBattleStats.hp)

		if chosenBattleStats.hp < 1 {
			cfg.Mux.Lock()
			fmt.Printf("Your %s has lost against %s. You lose it.\n\n", chosenPokemon.Name, rivalPokemon.Name)
			delete(cfg.CaughtPokemons, chosenPokemon.Name)
			cfg.CurrentOpponent = nil
			cfg.Mux.Unlock()
			break
		}

		fmt.Printf("Your attack #%d\n", i)

		rivalBattleStats.hp -= chosenBattleStats.attack
		fmt.Printf("Now your rival HP is %d\n\n", rivalBattleStats.hp)

		if rivalBattleStats.hp < 1 {
			fmt.Printf("Your %s has won against %s. Now, you have a new Pokemon. Check your Pokedex.\n\n", chosenPokemon.Name, rivalPokemon.Name)
			cfg.Mux.Lock()
			cfg.CaughtPokemons[rivalPokemon.Name] = *rivalPokemon
			cfg.CurrentOpponent = nil
			cfg.Mux.Unlock()
			break
		}
	}

	return nil
}
