package main

import (
	"github.com/maticardenas/matias-pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationArea *string // use use a pointer to nil to know if we are in the last location
	prevLocationArea *string // use use a pointer to nil to know if we are in the first location
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRepl(&cfg)
}
