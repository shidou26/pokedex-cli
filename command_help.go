package main

import (
	"fmt"
)

func commandHelp(cfg *config, argv []string) error {
	if len(argv) != 0 {
		return fmt.Errorf("usage: help")
	}
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for command, info := range cfg.registry {
		fmt.Printf("%v: %v\n", command, info.description)
	}
	return nil
}
