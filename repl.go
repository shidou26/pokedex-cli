package main

import (
	"bufio"
	"fmt"
	"github.com/shidou26/pokedex-cli/internal/pokecache"
	"os"
	"strings"
)

type config struct {
	registry             map[string]cliCommand
	locationAreaNext     string
	locationAreaPrevious string
	commandCache         pokecache.Cache
	pokedex              map[string]Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		isScanned := scanner.Scan()
		if !isScanned {
			break
		}
		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}
		command := input[0]
		commandInfo, ok := cfg.registry[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := commandInfo.callback(cfg, input[1:])
		if err != nil {
			fmt.Println(err)
		}
	}
}
