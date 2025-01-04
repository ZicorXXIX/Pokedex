package main

import (
    "os"
    "fmt"
    "net/http"
    "io"
    "encoding/json"

    // "github.com/mtslzr/pokeapi-go"
)

type cliCommand struct {
    name            string
    description     string
    callback        func() error
}
var commands map[string]cliCommand

var url = "https://pokeapi.co/api/v2/location-area/"
var prevUrl string
type LocationAreaResponse struct {
    Count int
    Next string
    Previous string
    Results []LocationArea
}

type LocationArea struct {
    Name string
    Url  string
}

func commandExit() error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp() error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    for _, cmd := range commands {
        fmt.Printf("\n%s: %s\n", cmd.name, cmd.description)
    }
    return nil
}

func commandMap() error {
    res, err := http.Get(url)
    if err != nil {
        return err
    }

    body, err := io.ReadAll(res.Body)
    if err != nil {
        return err
    }
    defer res.Body.Close()

    var data LocationAreaResponse
    if json.Valid([]byte(body)){
        json.Unmarshal([]byte(body), &data)
    }

    for _, area := range data.Results {
        fmt.Println(area.Name)
    }
    prevUrl = data.Previous
    url = data.Next
    return nil
}


func commandMapBack() error {
    if prevUrl == "" {
        fmt.Println("No previous location areas to display.")
    }
    res, err := http.Get(prevUrl)
    if err != nil {
        return err
    }

    body, err := io.ReadAll(res.Body)
    if err != nil {
        return err
    }
    defer res.Body.Close()

    var data LocationAreaResponse
    if json.Valid([]byte(body)){
        json.Unmarshal([]byte(body), &data)
    }

    for _, area := range data.Results {
        fmt.Println(area.Name)
    }
    url = data.Next
    prevUrl = data.Previous
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
