package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		formatted := formatInput(text)
		if len(formatted) == 0 {
			continue
		}
		commandName := formatted[0]

		commands := getCommands()
		command, ok := commands[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		command.cb()
	}
}

type cliCommand struct {
	name string
	desc string
	cb   func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			desc: "Prints the help menu",
			cb:   callbackHelp,
		},
		"exit": {
			name: "exit",
			desc: "Exit the Pokedex.",
			cb:   callbackExit,
		},
	}
}

func formatInput(str string) []string {
	low := strings.ToLower(str)
	words := strings.Fields(low)
	return words
}
