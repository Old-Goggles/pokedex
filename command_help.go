package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Print(
		"Welcome to the Pokedex!\n" +
			"Usage:\n" +
			"\n" +
			"help: Displays a help message\n" +
			"exit: Exit the Pokedex\n" +
			"map: Displays Locations in Pokemon World\n" +
			"mapb: isplays Previous Locations in Pokemon World\n" +
			"explore: Lists the Pokemon located in an area\n" +
			"catch: Attempts to catch a pokemon\n" +
			"inspect: Allow players to see details about a Pokemon\n" +
			"pokedex: Prints a list of all the names of the Pokemon the user has caught",
	)
	return nil
}
