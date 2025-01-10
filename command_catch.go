package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
    fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
    res, err := cfg.pokeapiClient.GetPokemonDetails(args[0])
    if err != nil {
        return err
    }
    xmax := 635
    xmin := 20
    normalized_base_experience := float64(res.BaseExperience -xmin)/float64(xmax - xmin) 
    // fmt.Println("BaseExperience: ", res.BaseExperience)
    // fmt.Println("normalized_base_experience:", normalized_base_experience *100)
    rand.Seed(time.Now().UnixNano())

    randomNumber := rand.Intn(101)
    // fmt.Println("randomNumber:", randomNumber)
    if int(normalized_base_experience *100) >= 100 {
        normalized_base_experience = 0.99
    }
    if randomNumber > int(normalized_base_experience *100) {
        fmt.Printf("You caught %s!\n", res.Name)
        cfg.pokedex[res.Name] = Pokemon{
            Name: res.Name,
            BaseExperience: res.BaseExperience,
        }
    } else {
        fmt.Printf("%s escaped!\n", res.Name)
    }

    return nil    
    
}
