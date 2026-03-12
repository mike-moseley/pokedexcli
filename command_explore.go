package main

import (
	"fmt"

	"github.com/mike-moseley/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, arg string) error {
	if arg == "" {
		fmt.Println("Please enter an area to explore.")
		return nil
	}
	pageURL := pokeapi.BaseURL + "/" + arg + "/"
	res, err := cfg.pokeapiClient.ExploreLocation(&pageURL)
	if err != nil {
		return err
	}

	if len(res.PokemonEncounters) == 0 {
		fmt.Println("No pokemon in this area.")
	}
	for _, poke := range res.PokemonEncounters {
		fmt.Println(poke.Pokemon.Name)
	}
	return nil
}
