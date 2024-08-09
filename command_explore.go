package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("explore command requires one location area name")
	}

	locationAreaName := args[0]

	resp, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Pokemon in %s:\n", locationAreaName)

	for _, area := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", area.Pokemon.Name)
	}
	return nil
}
