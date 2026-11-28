package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for command, info := range cfg.registry {
		fmt.Printf("%v: %v\n", command, info.description)
	}
	return nil
}
