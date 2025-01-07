package utils

import (
	"encoding/json"
	"fmt"
)

func processData(data []byte) error {
	var locationResponse LocationResponse
	if err := json.Unmarshal(data, &locationResponse); err != nil {
		return fmt.Errorf("sorry couldn't resolve the data: %v", err)
	}

	for _, value := range locationResponse.Results {
		fmt.Printf("%s\n", value.Name)
	}

	return nil
}
