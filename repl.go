package main

import (
	"bufio"
	"fmt"
	"github.com/arishimam/pokedexcli/internal/pokeapi"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	prevLocationsURL *string
	nextLocationsURL *string
	caughtPokemon    map[string]pokeapi.Pokemon
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
		"map": {
			name:        "map",
			description: "Displays next 20 locations, and so on.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 locations, and so on.",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Accepts a location argument and displays all pokemon in an area.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Gives the user a chance to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Allows user to view stats of a Pokemon they have caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Allows user to view names of all Pokemon they have caught",
			callback:    commandPokedex,
		},
	}
}

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		// input new bufio scanner
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := scanner.Text()
		output := cleanInput(input)

		if len(output) == 0 {
			continue
		}

		command := output[0]
		param := ""

		if len(output) > 1 {
			param = output[1]
		}

		// fmt.Println("PARAMETER: ", param)

		com, ok := supportedCommands[command]
		if ok {
			err := com.callback(cfg, param)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}

func cleanInput(text string) []string {
	lowerCase := strings.ToLower(text)
	output := strings.Fields(lowerCase)
	return output
}
