package utils

import "fmt"

func CommandHelp(commands map[string]CliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	commandOrder := []string{"help", "exit"}

	for _, cmdName := range commandOrder {
		cmd := commands[cmdName]
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}
