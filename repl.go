package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	registry             map[string]cliCommand
	locationAreaNext     string
	locationAreaPrevious string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Pokedex > ")
		isScanned := scanner.Scan()
		if !isScanned {
			break
		}
		input := cleanInput(scanner.Text())

		command := input[0]
		commandInfo, ok := cfg.registry[command]
		if !ok {
			fmt.Println("Uknown command")
			continue
		}

		err := commandInfo.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
