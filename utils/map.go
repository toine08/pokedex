package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var index int32

func CommandMap() error {
	if index == 0 {
		index = 20
	} else {
		index += 20
	}
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area?limit=%v", index)
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("sorry couldn't resolve the url: %v", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("couldn't get the data: %v", err)
	}
	var locationResponse LocationResponse

	if err := json.Unmarshal(data, &locationResponse); err != nil {
		return fmt.Errorf("sorry couldn't resolve the data: %v", err)
	}
	for _, value := range locationResponse.Results {
		fmt.Printf("%s\n", value.Name)
	}

	return nil
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
