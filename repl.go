package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var supportedCommands map[string]cliCommand

func init() {
	supportedCommands = map[string]cliCommand{
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

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		// input new bufio scanner
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		output := cleanInput(input)

		command := ""
		if len(output) > 0 {
			command = output[0]
		}

		if c, ok := supportedCommands[command]; ok {
			c.callback()

		} else {
			fmt.Println("Unknown command")
		}

	}
}

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	output := strings.Fields(lowerCase)

	return output

}

func commandExit() error {
	fmt.Println("\nClosing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil

}

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for k, v := range supportedCommands {
		fmt.Println(k + ":" + v.description)
	}

	fmt.Println()

	return nil
}
