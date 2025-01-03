package main

import "fmt"
func main() {
    fmt.Println("Welcome to the Pokedex CLI!")
	fmt.Println("Type 'help' for a list of commands or 'exit' to quit.")

    initCommands()

    startRepl()
}
