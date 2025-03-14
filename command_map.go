package main

import (
	"errors"
	"fmt"
)

type locationAreas struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMapf(cfg *config, args ...string) error {
	locationsResp, err := cfg.pokeApiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.previousLocationsURL = locationsResp.Previous

    for _, loc := range locationsResp.Results {
        fmt.Println(loc.Name)
    }
    fmt.Println()
    return nil
}

func commandMapb(cfg *config, args ...string) error {
    if cfg.previousLocationsURL == nil {
        return errors.New("you're on the first page")
    }

    locationsResp, err := cfg.pokeApiClient.ListLocations(cfg.previousLocationsURL)
    if err != nil {
        return err
    }

    cfg.nextLocationsURL = locationsResp.Next
    cfg.previousLocationsURL = locationsResp.Previous

    for _, loc := range locationsResp.Results {
        fmt.Println(loc.Name)
    }
    fmt.Println()
    return nil
}
