package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
			Callback: func(args ...string) error {
				if len(args) > 0 {
					return fmt.Errorf("this command doesn't accept any arguments")
				}
				return utils.CommandExit()
			},
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback: func(args ...string) error {
				if len(args) > 0 {
					return fmt.Errorf("this command doesn't accept any arguments")
				}
				return callBackHelp()
			},
		},
		"map": {
			Name:        "map",
			Description: "Displays 20 names of locations",
			Callback: func(args ...string) error {
				if len(args) > 0 {
					return fmt.Errorf("this command doesn't accept any arguments")
				}
				return utils.CommandMap()
			},
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the 20 previous names of locations",
			Callback: func(args ...string) error {
				if len(args) > 0 {
					return fmt.Errorf("this command doesn't accept any arguments")
				}
				return utils.CommandMapB()
			},
		},
		"explore": {
			Name:        "explore",
			Description: "Explore an area",
			Callback: func(args ...string) error {
				if len(args) < 1 {
					return fmt.Errorf("please provide a zone to explore")
				}
				return utils.CommandExplore(args[0])
			},
		},
		"catch": {
			Name:        "catch",
			Description: "catch a pokemon",
			Callback: func(args ...string) error {
				if len(args) < 1 {
					return fmt.Errorf("please provide a pokemon to catch")
				}
				return utils.CommandCatch(args[0])
			},
		},
		"inspect": {
			Name:        "inpsect",
			Description: "inspect a pokemon",
			Callback: func(args ...string) error {
				if len(args) < 1 {
					return fmt.Errorf("please provide a pokemon to inspect")
				}
				return utils.CommandInspect(args[0])
			},
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "inspect your pokedex",
			Callback: func(args ...string) error {
				if len(args) > 0 {
					return fmt.Errorf("this command doesn't accept any arguments")
				}
				return utils.CommandPokedex()
			},
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()

		// Split input into command and arguments
		inputParts := strings.Fields(input)
		if len(inputParts) == 0 {
			continue // Skip empty input
		}

		command := inputParts[0] // First part is the command
		args := inputParts[1:]   // Remaining parts are arguments (if any)

		if cmd, exists := commands[command]; exists {
			if err := cmd.Callback(args...); err != nil { // Pass args to the callback
				fmt.Println("Error: ", err)
			}
		} else {
			fmt.Println("Unknown command: ", command)
		}
	}
}
