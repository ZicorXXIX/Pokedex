package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ZicorXXIX/pokedex/internal/pokeapi"
)

type config struct {
    pokeapiClient pokeapi.Client
    pokedex       map[string]pokeapi.Pokemon
    nextPage      *string
    prevPage      *string
}

type cliCommand struct {
    name            string
    description     string
    callback        func(*config, ...string) error
}

func startRepl(cfg *config){
    reader := bufio.NewScanner(os.Stdin)
    for{
        fmt.Print("Pokedex>")
        if !reader.Scan() {
            continue
        }
        input := reader.Text()

        parts := strings.Fields(input)
        if len(parts) == 0 {
            continue
        }

        cmdName := parts[0]
        args := parts[1:]

        cmd, found := commands[cmdName]

        if found {
            if err:= cmd.callback(cfg, args...); err != nil{
                fmt.Printf("Error: %v \n",err)
            }
        }else {
            fmt.Println("Unknown Command")
        }
    }
  // How i would have implemented without instructions from bood.dev
  //   reader:= bufio.NewReader(os.Stdin)
  // //comma ok || error ok syntax
  // input, _ := reader.ReadString('\n')

}

func cleanInput(text string) []string {
    fields := strings.Fields(text)

    for i, s := range fields {
        fields[i] = strings.ToLower(s)
    }

    return fields
}

