package main

import (
	"time"

	"github.com/mike-moseley/pokedexcli/internal/pokeapi"
)

func main() {
	config := config{
		pokeapiClient: pokeapi.NewClient(500 * time.Millisecond),
		pokedex:       make(map[string]pokeapi.Pokemon),
	}
	startRepl(&config)
}
