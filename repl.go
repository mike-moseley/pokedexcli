package main

import "strings"
import "os"
import "bufio"
import "fmt"
import "github.com/mike-moseley/pokedexcli/internal/pokeapi"

func startRepl(c *config) {
	for k := range commands {
		fmt.Printf("%s: %s\n", k, commands[k].description)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		input := scanner.Text()
		input_slice := cleanInput(input)
		command := input_slice[0]
		var argument string
		if len(input_slice) > 1 {
			argument = input_slice[1]
		}
		_, ok := commands[command]
		if ok {
			if err := commands[command].callback(c, argument); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Command not found")
		}
	}
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	text_slice := strings.Split(text, " ")

	return text_slice
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Help using the Pokedex",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "Show next map list",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Show previous map list",
		callback:    commandMapb,
	},
	"explore": {
		name:        "explore",
		description: "Explore the pokemon in an area",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "Catch a pokemon",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect",
		description: "Inspect a pokemon from your pokedex",
		callback:    commandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "View pokedex",
		callback:    commandPokedex,
	},
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	pokeapiClient        pokeapi.Client
	pokedex              map[string]pokeapi.Pokemon
	nextLocationsURL     *string
	previousLocationsURL *string
}
