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
	cb   func(*config, ...string) error
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
		args := []string{}
		if len(formatted) > 1 {
			args = formatted[1:]
		}

		commands := getCommands()

		command, ok := commands[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		err := command.cb(cfg, args...)
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
		"explore": {
			name: "explore {location_area}",
			desc: "List pokemons on a given area",
			cb:   callbackExplore,
		},
		"catch": {
			name: "catch {pokemon_name}",
			desc: "Try to catch a given pokemon",
			cb:   callbackCatch,
		},
		"inspect": {
			name: "inspect {pokemon_name}",
			desc: "Inspect a caught pokemon",
			cb:   callbackInspect,
		},
		"pokedex": {
			name: "pokedex",
			desc: "List every caught pokemon",
			cb:   callbackPokedex,
		},
		"exit": {
			name: "exit",
			desc: "Exit the Pokedex",
			cb:   callbackExit,
		},
	}
}

func formatInput(str string) []string {
	low := strings.ToLower(str)
	words := strings.Fields(low)
	return words
}
