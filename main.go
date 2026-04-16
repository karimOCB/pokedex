package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	registry := getCommand()

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		line := scanner.Text()
		words := cleanInput(line)

		val, ok := registry[words[0]]
		if ok {
			val.callback()
		} else {
			fmt.Println("Unknown command")
		}
	}
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
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	registry := getCommand()

	for key, value := range registry {
		fmt.Printf("%s: %s\n", key, value.description)
	}
	return nil
}
