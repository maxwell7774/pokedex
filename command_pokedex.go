package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
    fmt.Println("Your Pokedex:")
    for key := range cfg.caughtPokemon {
        fmt.Printf("  - %s\n", key)
    }
    fmt.Println()
    return nil
}
