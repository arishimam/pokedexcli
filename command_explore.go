package main

import "fmt"

func commandExplore(cfg *config, param string) error {
	areaDetails, err := cfg.pokeapiClient.ListPokemon(param)
	if err != nil {
		return err
	}

	for _, v := range areaDetails.PokemonEncounters {
		fmt.Println(v.Pokemon.Name)
	}

	return nil
}
