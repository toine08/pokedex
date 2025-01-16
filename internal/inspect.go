package utils

import (
	"fmt"
)

func CommandInspect(pokemon string) error {
	if inspectedPokemon, caught := PokemonCatched[pokemon]; caught {
		fmt.Printf("Name: %s\nHeight: %d\nStats:\n", inspectedPokemon.Name, inspectedPokemon.Height)
		for _, stat := range inspectedPokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types: \n")
		for _, types := range inspectedPokemon.Types {
			fmt.Printf("  -%s\n", types.Type.Name)
		}

	} else {
		// pokemon was not found
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
