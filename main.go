package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/toine08/pokedexcli/utils"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the program",
			callback:    utils.CommandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		if cmd, exists := commands[input]; exists {
			if err := cmd.callback(); err != nil {
				fmt.Println("Error: ", err)
			}
		} else {
			fmt.Println("Unknown command: ", input)
		}
	}
	//slicedInput := utils.CleanInput(input)
}
