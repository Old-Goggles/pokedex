package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("NO Pokemon found")
	}
	pokemon := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	p, err := cfg.pokeClient.GetPokemon(pokemon)
	if err != nil {
		return fmt.Errorf("failed to find Pokemon: %w", err)
	}

	chance := 5

	if p.BaseExperience < 50 {
		chance = 80
	} else if p.BaseExperience < 100 {
		chance = 60
	} else {
		chance = 40
	}

	catchRoll := rand.Intn(100)
	if catchRoll < chance {
		fmt.Printf("%s was caught!\n", p.Name)
		fmt.Printf("You may now inspect it with the inspect command")
		cfg.pokedex[p.Name] = p
	} else {
		fmt.Printf("%s escaped!\n", p.Name)
	}

	return nil
}
