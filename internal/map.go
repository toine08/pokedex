package utils

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	index int32
	cache *Cache // Global cache instance
)

func init() {
	// Initialize the global cache with a 5-second interval for reaping
	cache = NewCache(5 * time.Second)
}

func CommandMap() error {
	if index == 0 {
		index = 20
	} else {
		index += 20
	}

	baseUrl := "https://pokeapi.co/api/v2/location-area"
	url := fmt.Sprintf("%s?limit=%v", baseUrl, index)

	if data, ok := cache.Get(url); ok {
		return processData(data)
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

	return processData(data)
}
func CommandMapB() error {
	if index < 20 {
		fmt.Print("you're on the first page")
	} else {
		index -= 40
		CommandMap()
	}
	return nil
}
