package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
)

var PokemonCatched = make(map[string]Pokemon)

func CommandCatch(pokemon string) error {
	baseUrl := "https://pokeapi.co/api/v2/pokemon/"
	url := fmt.Sprintf("%s/%s", baseUrl, pokemon)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	if data, ok := cache.Get(url); ok {
		var catchedPokemon Pokemon
		if err := json.Unmarshal(data, &catchedPokemon); err != nil {
			return fmt.Errorf("sorry couldn't resolve the data: %v", err)
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("sorry couldn't resolve the url: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("couldn't get the data: %v", err)
	}
	cache.Add(url, data)

	var catchedPokemon Pokemon
	if err := json.Unmarshal(data, &catchedPokemon); err != nil {
		return fmt.Errorf("sorry couldn't resolve the data: %v", err)
	}
	chanceToCatch := rand.Intn(100)
	threshold := int(math.Max(10, 100-float64(catchedPokemon.Level/2)))
	//fmt.Printf("this is chance to catch: %v, and this is treshold: %v\n", chanceToCatch, threshold) // adjust this formula to make it reasonably catchable
	if threshold > chanceToCatch {
		PokemonCatched[pokemon] = catchedPokemon
		fmt.Printf("%s was caught!\n", pokemon)
		fmt.Printf("You may now inspect it with the inspect command.\n")
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}
