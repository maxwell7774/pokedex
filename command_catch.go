package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
    if len(args) != 1 {
        return errors.New("A pokemon name must be provided")
    }

    name := args[0]

    pokemonResp, err := cfg.pokeApiClient.PokemonDetails(name)
    if err != nil {
        return err
    }

    fmt.Printf("Throwing a Pokeball at %s...\n", name)
    caught := rand.Int32N(int32(pokemonResp.BaseExperience)) <= 40
    if caught {
        cfg.caughtPokemon[pokemonResp.Name] = pokemonResp
        fmt.Printf("You caught %s!\n", pokemonResp.Name)
    } else {
        fmt.Printf("Failed to catch %s... Please try again.\n", pokemonResp.Name)
    }
    fmt.Println()

    return nil
}
