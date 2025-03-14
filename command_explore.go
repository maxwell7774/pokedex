package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("A location name must be provided")
	}

	locationName := args[0]

	fmt.Printf("Exploring %s...\n", locationName)
	locationResp, err := cfg.pokeApiClient.LocationDetails(locationName)
	if err != nil {
		return err
	}

	fmt.Println("Found pokemon:")
	for _, encounter := range locationResp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
    fmt.Println()

	return nil
}
