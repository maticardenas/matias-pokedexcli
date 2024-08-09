package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("Here are the pokemons you have caught:")
	for _, pokemon := range cfg.caughtPokemons {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}
