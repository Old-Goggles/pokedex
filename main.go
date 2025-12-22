package main

import (
	"pokedex/internal/pokeapi"
	"pokedex/internal/pokecache"
	"time"
)

func main() {
	cache := pokecache.NewCache(5 * time.Minute)
	client := pokeapi.NewClient(5*time.Second, cache)

	cfg := config{
		pokeClient: client,
		pokedex:    make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}
