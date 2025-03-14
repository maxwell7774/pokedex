package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("A pokemon name must be provided")
	}

	name := args[0]
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("Pokemon not caught yet...")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
    for _, pokemonType := range pokemon.Types {
		fmt.Printf("  -%s\n", pokemonType.Type.Name)
    }
	fmt.Println()
	return nil
}
