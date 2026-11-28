package main

import (
	"fmt"
)

func commandPokedex(cfg *config, argv []string) error {
	if len(argv) != 0 {
		return fmt.Errorf("usage: pokedex")
	}

	fmt.Println("Your pokedex:")
	for key := range cfg.pokedex {
		fmt.Printf(" - %v\n", key)
	}

	return nil
}
