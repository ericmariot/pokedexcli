package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("no information about %s", pokemon.Name)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Number: %v\n", pokemon.ID)
	fmt.Println("Types:")
	for _, types := range pokemon.Types {
		fmt.Printf(" - %s\n", types.Type.Name)
	}
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Abilities:")
	for _, ability := range pokemon.Abilities {
		fmt.Printf(" - %s\n", ability.Ability.Name)
	}
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	return nil
}
