package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}

	cfg.nextLocationURL = resp.Next
	cfg.prevLocationURL = resp.Previous

	return nil
}

func callbackMapb(cfg *config, args ...string) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you're on the first page")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}

	cfg.nextLocationURL = resp.Next
	cfg.prevLocationURL = resp.Previous

	return nil
}
