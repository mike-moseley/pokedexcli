package main

import (
	"time"

	"github.com/mike-moseley/pokedexcli/internal/pokeapi"
)

func main() {
	config := config{
		pokeapiClient: pokeapi.NewClient(500 * time.Millisecond),
	}
	startRepl(&config)
}
