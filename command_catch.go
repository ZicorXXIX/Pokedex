package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
    if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

    res, err := cfg.pokeapiClient.GetPokemonDetails(args[0])
    if err != nil {
        return err
    }
    fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
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
        cfg.pokedex[res.Name] = res  
    } else {
        fmt.Printf("%s escaped!\n", res.Name)
    }

    return nil    
    
}
