package main

import (
	"bufio"
	"fmt"
	"os"

	utils "github.com/toine08/pokedexcli/internal"
)

func main() {

	var commands map[string]utils.CliCommand
	callBackHelp := func() error {
		return utils.CommandHelp(commands)
	}

	commands = map[string]utils.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the program",
			Callback:    utils.CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    callBackHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays 20 names of locations",
			Callback:    utils.CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the 20 previous names of locations",
			Callback:    utils.CommandMapB,
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
			if err := cmd.Callback(); err != nil {
				fmt.Println("Error: ", err)
			}
		} else {
			fmt.Println("Unknown command: ", input)
		}
	}
	//slicedInput := utils.CleanInput(input)
}
