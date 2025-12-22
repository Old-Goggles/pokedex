package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeClient pokeapi.Client
	next       string
	previous   string
	pokedex    map[string]pokeapi.Pokemon
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Displays Locations in Pokemon World",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays Previous Locations in Pokemon World",
		callback:    commandMapB,
	},
	"explore": {
		name:        "explore",
		description: "Lists the Pokemon located in an area",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "Attempts to catch a pokemon",
		callback:    commandCatch,
	},
}

func cleanInput(text string) []string {
	trim := strings.TrimSpace(text)
	lower := strings.ToLower(trim)
	field := strings.Fields(lower)

	return field
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := commands[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
