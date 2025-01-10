package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
    for _, pokemon := range cfg.pokedex {
        fmt.Println("Your Pokedex")
        fmt.Printf(" - %s\n", pokemon.Name)
    }
    return nil
}
