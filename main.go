package main

import (
	"github.com/arishimam/pokedexcli/internal/cache"
	"github.com/arishimam/pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(50 * time.Second)
	C := cache.NewCache(5 * time.Second)

	cfg := &config{
		pokeapiClient:    pokeClient,
		prevLocationsURL: nil,
		nextLocationsURL: nil,
		cache:            C,
	}

	startRepl(cfg)
}
