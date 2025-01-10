package main


var commands map[string]cliCommand

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
        "explore": {
            name: "explore",
            description: "Explore the location area selected.",
            callback: commandExplore,
        },
        "catch": {
            name: "catch",
            description: "Catch a Pokemon",
            callback: commandCatch,
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
