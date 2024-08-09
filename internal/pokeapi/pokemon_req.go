package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (PokemonResponse, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	// checks if the URL is already in the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("Cache hit")
		pokemonResp := PokemonResponse{}
		err := json.Unmarshal(dat, &pokemonResp)

		if err != nil {
			return PokemonResponse{}, err
		}

		return pokemonResp, nil
	}

	// builds the request
	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return PokemonResponse{}, err
	}

	// actually performs the request through the client
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer resp.Body.Close() // close the resp object when finished.

	if resp.StatusCode > 399 {
		return PokemonResponse{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	pokemonResp := PokemonResponse{}
	err = json.Unmarshal(dat, &pokemonResp)

	if err != nil {
		return PokemonResponse{}, err
	}

	c.cache.Add(fullURL, dat)

	return pokemonResp, nil
}
