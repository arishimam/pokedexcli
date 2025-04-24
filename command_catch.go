package main

import (
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(cfg *config, param string) error {
	fmt.Printf("Throwing a Pokeball at %v...\n", param)
	pokeInfo, err := cfg.pokeapiClient.PokeInfo(param)
	if err != nil {
		return err
	}

	if randomOutcome(pokeInfo.BaseExperience) {
		fmt.Printf("%v was caught!\n", param)
		cfg.caughtPokemon[pokeInfo.Name] = pokeInfo
		//fmt.Println("Added to pokedex!")
	} else {
		fmt.Printf("%v escaped!\n", param)
	}

	return nil
}

func randomOutcome(baseExp int) bool {
	decayRate := 0.004 // tune this to control how fast probability drops
	probability := math.Exp(-decayRate * float64(baseExp))

	randVal := rand.Float64()
	//fmt.Println("baseExp = ", baseExp)
	//fmt.Println("probability= ", probability)
	//fmt.Println("randVal = ", randVal)

	return randVal < probability
}
