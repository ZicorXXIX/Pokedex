package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
    res, err := cfg.pokeapiClient.GetPokemonEncounters(args[0])         
    if err != nil {
        return err
    }
    fmt.Printf("Exploring %s...\n", args[0])
    fmt.Println("Found Pokemon:")
    for _, pokemon := range res.PokemonEncounters {
        fmt.Printf(" - %s\n",pokemon.Pokemon.Name)
    }
    return nil
}
