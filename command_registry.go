package main

import (
    "os"
    "fmt"
)

type cliCommand struct {
    name            string
    description     string
    callback        func() error
}
var commands map[string]cliCommand

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
