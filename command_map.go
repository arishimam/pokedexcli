package main

import (
	"fmt"
)

func commandMap(cfg *config) error {

	pokeAreas, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = pokeAreas.Next
	cfg.prevLocationsURL = pokeAreas.Previous

	for _, a := range pokeAreas.Results {
		fmt.Println(a.Name)
	}

	return nil
}

func commandMapB(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	cfg.nextLocationsURL = cfg.prevLocationsURL
	return commandMap(cfg)

}
