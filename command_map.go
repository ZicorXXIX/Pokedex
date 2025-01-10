package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
    res, err := cfg.pokeapiClient.GetLocations(cfg.nextPage)
    if err != nil {
        return err
    }


    for _, area := range res.Results {
        fmt.Println(area.Name)
    }
    cfg.prevPage= res.Previous
    cfg.nextPage= res.Next
    return nil
}


func commandMapBack(cfg *config, args ...string) error {
    if cfg.prevPage == nil {
        fmt.Println("No previous location areas to display.")
        return errors.New("You're on first page")
    }
    res, err := cfg.pokeapiClient.GetLocations(cfg.prevPage)
    if err != nil {
        return err
    }

    for _, area := range res.Results {
        fmt.Println(area.Name)
    }
    cfg.nextPage = res.Next
    cfg.prevPage = res.Previous
    return nil
}


