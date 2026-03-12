package main

import (
	"fmt"
)

func commandMap(cfg *config, arg string) error {
	res, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = res.Next
	cfg.previousLocationsURL = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, arg string) error {
	res, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL)

	if err != nil {
		return err
	}
	cfg.nextLocationsURL = res.Next
	cfg.previousLocationsURL = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
