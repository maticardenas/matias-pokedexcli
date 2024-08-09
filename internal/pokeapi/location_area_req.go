package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	// checks if the URL is already in the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("Cache hit")
		locationAreasResp := LocationAreasResponse{}
		err := json.Unmarshal(dat, &locationAreasResp)

		if err != nil {
			return LocationAreasResponse{}, err
		}

		return locationAreasResp, nil
	}

	// builds the request
	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	// actually performs the request through the client
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close() // close the resp object when finished.

	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreasResp := LocationAreasResponse{}
	err = json.Unmarshal(dat, &locationAreasResp)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationAreasResp, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationAreaResponse, error) {
	endpoint := "/location-area" + "/" + locationAreaName
	fullURL := baseURL + endpoint

	// checks if the URL is already in the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("Cache hit")
		locationAreaResp := LocationAreaResponse{}
		err := json.Unmarshal(dat, &locationAreaResp)

		if err != nil {
			return LocationAreaResponse{}, err
		}

		return locationAreaResp, nil
	}

	// builds the request
	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreaResponse{}, err
	}

	// actually performs the request through the client
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close() // close the resp object when finished.

	if resp.StatusCode > 399 {
		return LocationAreaResponse{}, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationAreasResp := LocationAreaResponse{}
	err = json.Unmarshal(dat, &locationAreasResp)

	if err != nil {
		return LocationAreaResponse{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationAreasResp, nil
}
