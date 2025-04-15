package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range supportedCommands {
		fmt.Println(cmd.name + ": " + cmd.description)
	}

	fmt.Println()

	return nil
}
