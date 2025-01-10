package main


var commands map[string]cliCommand

func initCommands() {
    commands = map[string]cliCommand{
       "help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempt to catch a pokemon",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapBack,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
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
