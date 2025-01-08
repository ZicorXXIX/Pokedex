package main

import (
	"errors"
	"fmt"
)
var commands map[string]cliCommand

func commandMap(cfg *config) error {
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


func commandMapBack(cfg *config) error {
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


func initCommands() {
    commands = map[string]cliCommand{
        "exit": {
            name: "exit",
            description: "Exit the Pokedex",
            callback: commandExit,
        },
        "help": {
            name: "help",
            description: "Displays help message",
            callback: commandHelp,
        },
        "map": {
            name: "map",
            description: "It displays the names of 20 location areas in the Pokemon world.",
            callback: commandMap,
        },
        "mapb": {
            name: "mapb(map back)",
            description:  "Displays previous 20 location areas in the Pokemon world.",
            callback: commandMapBack,
        },
    }
}

//Better way to solve circular dependency during initialization (Boot Dev's way)
// func getCommands() map[string]cliCommand {
// 	return map[string]cliCommand{
// 		"help": {
// 			name:        "help",
// 			description: "Displays a help message",
// 			callback:    commandHelp,
// 		},
// 		"exit": {
// 			name:        "exit",
// 			description: "Exit the Pokedex",
// 			callback:    commandExit,
// 		},
// 	}
// }
