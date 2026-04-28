package main

import "fmt"

func commandRun(cfg *config, args ...string) error {

	cfg.Mux.Lock()
	if cfg.CurrentOpponent == nil {
		cfg.Mux.Unlock()
		fmt.Fprintf(cfg.RlInstance,"There is no wild Pokemon to run from right now!\n")
		return nil
	}

	opponentName := cfg.CurrentOpponent.Name
	cfg.CurrentOpponent = nil 
	cfg.Mux.Unlock()
	fmt.Fprintf(cfg.RlInstance, "You have succesfully run from: %s\n", opponentName)

	return nil
}