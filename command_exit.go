package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, argv []string) error {
	if len(argv) != 0 {
		return fmt.Errorf("usage: exit")
	}
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
