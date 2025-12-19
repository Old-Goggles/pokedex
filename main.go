package main

import (
	"time"

	"pokedex/internal/pokeapi"
	"pokedex/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Minute)
	client := pokeapi.NewClient(5*time.Second, cache)

	cfg := config{
		pokeClient: client,
	}

	startRepl(&cfg)
}
