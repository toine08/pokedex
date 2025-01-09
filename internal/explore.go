package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CommandExplore(zone string) error {
	baseUrl := "https://pokeapi.co/api/v2/location-area"
	url := fmt.Sprintf("%s/%s", baseUrl, zone)

	fmt.Printf("Exploring %s...\n", zone)

	if data, ok := cache.Get(url); ok {
		var locationArea LocationArea
		if err := json.Unmarshal(data, &locationArea); err != nil {
			return fmt.Errorf("sorry couldn't resolve the data: %v", err)
		}

		for _, value := range locationArea.PokemonEncounters {
			fmt.Printf("- %s\n", value.Pokemon.Name)
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

	var locationArea LocationArea
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return fmt.Errorf("sorry couldn't resolve the data: %v", err)
	}
	for _, value := range locationArea.PokemonEncounters {
		fmt.Printf("- %s\n", value.Pokemon.Name)
	}
	return nil
}
