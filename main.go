package main

import (
	"time"

	"pokedex/internal/pokeapi"
)

var pokeClient = pokeapi.NewClient(5 * time.Second)

func main() {
	startRepl()
}
