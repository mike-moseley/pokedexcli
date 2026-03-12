package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, arg string) error {
	if arg == "" {
		fmt.Println("Please enter a pokemon.")
		return nil
	}
	pageURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", arg)
	pokemon, err := cfg.pokeapiClient.GetPokemonInfo(&pageURL)
	if err != nil {
		return fmt.Errorf("Error getting pokemon info: %w", err)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if rand.Intn(700) > pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.pokedex[pokemon.Name] = pokemon
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
