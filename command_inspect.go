package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("catch command requires one pokemon name")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemons[pokemonName]
	if !ok {
		return errors.New("you haven't caught this pokemon yet")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Base Experience: %d\n", pokemon.BaseExperience)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	for _, stat := range pokemon.Stats {
		fmt.Printf("Stat: %s, Value: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	for _, typ := range pokemon.Types {
		fmt.Printf("Type: %s\n", typ.Type.Name)
	}

	return nil
}
