package main

import (
	"strings"
)

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	for i := range words {
		words[i] = strings.Trim(words[i], " ")
	}
	return words
}
