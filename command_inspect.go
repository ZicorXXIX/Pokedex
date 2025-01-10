package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
    pokemonName := args[0]

    val, ok := cfg.pokedex[pokemonName]
    if ok {
        fmt.Println("Name: ", val.Name)
        fmt.Println("Height: ", val.Height)
        fmt.Println("Weight: ", val.Weight)
        fmt.Println("Stats:")
        for _, stats := range val.Stats {
            fmt.Printf(" -%s: %d\n", stats.Stat.Name, stats.BaseStat)
        }
        fmt.Println("Types:")
        for _, value := range val.Types {
            fmt.Printf(" - %s\n", value.Type.Name)
        }
    } else {
        return errors.New("you have not caught that pokemon")
    }

    return nil
}
