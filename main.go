package main

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
				name:        "map",
				description: "Show previous location areas",
				callback:    commandMapBack,
			},
		},
		locationAreaNext:     "https://pokeapi.co/api/v2/location-area",
		locationAreaPrevious: "",
	}
	startRepl(cfg)

}
