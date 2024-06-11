package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	desc string
	cb   func(*config) error
}

func startRepl(cfg *config) {
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

		err := command.cb(cfg)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println()
	}
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
			desc: "Exit the Pokedex",
			cb:   callbackExit,
		},
		"map": {
			name: "map",
			desc: "List 20 location areas",
			cb:   callbackMap,
		},
		"mapb": {
			name: "mapb",
			desc: "List 20 previous location areas",
			cb:   callbackMapb,
		},
	}
}

func formatInput(str string) []string {
	low := strings.ToLower(str)
	words := strings.Fields(low)
	return words
}
