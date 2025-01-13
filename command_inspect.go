package main
import (
	"errors"
	"fmt"

    "github.com/ZicorXXIX/Image2ASCII/pkg/convert"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, ok := cfg.pokedex[name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

    //For image to ascii
    converter := convert.NewImageConverter()
    options := convert.DefaultOptions
    options.Width = 20
    options.Height = 20
    img := converter.ImageUrlToAsciiString(pokemon.Sprites.FrontDefault, options)

    fmt.Println(img)
	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Println("  -", typeInfo.Type.Name)
	}
	return nil
}

