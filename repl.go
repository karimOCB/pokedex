package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/karimOCB/pokedex/internal/pokeapi"
)

type config struct {
	Next     *string
	Previous *string
	Client   *pokeapi.Client
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	registry := getCommand()

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		cmd, ok := registry[words[0]]
		if ok {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}

func getCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 next locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays 20 previous locations",
			callback:    commandMapb,
		},
	}
}
