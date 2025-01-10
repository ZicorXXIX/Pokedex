package main
import (
	"fmt"
	"time"

	"github.com/ZicorXXIX/pokedex/internal/pokeapi"
)
func main() {
    fmt.Println("Welcome to the Pokedex CLI!")
	fmt.Println("Type 'help' for a list of commands or 'exit' to quit.")

    initCommands()
    pokeClient := pokeapi.NewClient(5 *time.Second, 5 *time.Minute)
    cfg := &config{
        pokeapiClient: pokeClient,
        pokedex: make(map[string]Pokemon),
    }

    startRepl(cfg)
}
