package main

import (
	"fmt"
)

func commandMap(c *config) error {
	res, err := c.pokeapiClient.ListLocations(c.nextLocationsURL)
	if err != nil {
		return err
	}
	c.nextLocationsURL = res.Next
	c.previousLocationsURL = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(c *config) error {
	res, err := c.pokeapiClient.ListLocations(c.previousLocationsURL)

	if err != nil {
		return err
	}
	c.nextLocationsURL = res.Next
	c.previousLocationsURL = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
