package main

import (
	"fmt"
	//"math/rand"
)

func commandCatch(cfg *config, param string) error {
	fmt.Printf("Throwing a Pokeball at %v...\n", param)
	pokeInfo, err := cfg.pokeapiClient.PokeInfo(param)
	if err != nil {
		return err
	}

	fmt.Println(pokeInfo)

	return nil

}
