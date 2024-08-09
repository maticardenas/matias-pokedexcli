package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("catch command requires one pokemon name")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Pokemon in %s:\n", pokemonName)
	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Println("Trying to catch the pokemon..." + string(pokemon.BaseExperience) + " - " + pokemonName)

	if randNum > threshold {
		return fmt.Errorf("you failed to catch %s!", pokemon.Name)
	}

	cfg.caughtPokemons[pokemonName] = pokemon
	fmt.Printf("You caught %s!\n", pokemon.Name)
	return nil
}
