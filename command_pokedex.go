package main

import (
	"fmt"
)

func commandPokedex(cfg *config, param string) error {
	fmt.Println("Your pokedex:")
	for _, poke := range cfg.caughtPokemon {
		fmt.Printf("  - %s\n", poke.Name)
	}

	return nil
}
