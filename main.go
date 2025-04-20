package main

import (
	"github.com/arishimam/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Second)

	cfg := &config{
		pokeapiClient:    pokeClient,
		prevLocationsURL: nil,
		nextLocationsURL: nil,
	}

	startRepl(cfg)
}
