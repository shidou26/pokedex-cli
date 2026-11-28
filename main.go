package main

import (
	"github.com/shidou26/pokedex-cli/internal/pokecache"
	"time"
)

func main() {
	cfg := &config{
		registry: map[string]cliCommand{
			"help": {
				name:        "help",
				description: "Show available commands",
				callback:    commandHelp,
			},
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"map": {
				name:        "map",
				description: "Show next location areas",
				callback:    commandMap,
			},
			"mapb": {
				name:        "mapb",
				description: "Show previous location areas",
				callback:    commandMapBack,
			},
			"explore": {
				name:        "explore",
				description: "Show list of Pokemon found in an area",
				callback:    commandExplore,
			},
			"catch": {
				name:        "catch",
				description: "Catch a Pokemon!",
				callback:    commandCatch,
			},
			"inspect": {
				name:        "inspect",
				description: "Inspect a Pokemon",
				callback:    commandInspect,
			},
			"pokedex": {
				name:        "pokedex",
				description: "Show your Pokedex",
				callback:    commandPokedex,
			},
		},
		locationAreaNext:     "https://pokeapi.co/api/v2/location-area",
		locationAreaPrevious: "",
		commandCache:         *pokecache.NewCache(5 * time.Minute),
		pokedex:              make(map[string]Pokemon),
	}
	startRepl(cfg)
}
