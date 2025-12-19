package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide location name")
	}
	location := args[0]
	fmt.Printf("Exploring %s...\n", location)

	loc, err := cfg.pokeClient.GetLocation(location)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range loc.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
