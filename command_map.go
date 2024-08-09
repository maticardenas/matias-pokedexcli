package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationArea)
	if err != nil {
		log.Fatal(err)
	}

	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationArea = resp.Next
	cfg.prevLocationArea = resp.Previous
	return nil
}

func callbackMapBack(cfg *config, args ...string) error {
	if cfg.prevLocationArea == nil {
		return errors.New("already in first location area")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationArea)
	if err != nil {
		log.Fatal(err)
	}

	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationArea = resp.Next
	cfg.prevLocationArea = resp.Previous
	return nil
}
