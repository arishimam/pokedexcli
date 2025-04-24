package main

import (
	"fmt"
)

func commandInspect(cfg *config, param string) error {

	if pokemon, exists := cfg.caughtPokemon[param]; exists {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, val := range pokemon.Stats {
			fmt.Printf("  -%s: %v\n", val.Stat.Name, val.BaseStat)
		}
		fmt.Println("Types:")
		for _, val := range pokemon.Types {
			fmt.Printf("  - %s\n", val.Type.Name)
		}

		return nil

	}
	fmt.Println("you have not caught that pokemon")

	return nil
}

