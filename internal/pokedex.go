package utils

import (
	"fmt"
)

func CommandPokedex() error {
	fmt.Printf("Your pokedex:\n")
	for pokemon := range PokemonCatched {
		fmt.Printf("  - %s\n", pokemon)
	}

	return nil
}
