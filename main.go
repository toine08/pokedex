package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/toine08/pokedexcli/utils"
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
