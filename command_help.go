package main

import (
    "fmt"
)

func commandHelp(cfg *config, args ...string) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    for _, cmd := range commands {
        fmt.Printf("\n%s: %s\n", cmd.name, cmd.description)
    }
    return nil
}
