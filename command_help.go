package main

import (
	"fmt"
)

func callbackHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%v: %v\n", command.name, command.desc)
	}

	fmt.Println()
	return nil
}
