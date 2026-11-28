package main 

import (
	"fmt"
)

func commandInspect(cfg *config, argv []string) error {
	if len(argv) != 1 {
		return fmt.Errorf("usage: inspect <pokemon-name>")
	}

	info, ok := cfg.pokedex[argv[0]]
	if !ok {
		return fmt.Errorf("you have not caught %v", argv[0])
	}

	fmt.Printf("Name: %v\n", info.Name)
	fmt.Printf("Height: %v\nWeight: %v\n", info.Height, info.Weight)
	fmt.Println("Stats:")
	for _, stat := range info.Stats {
		fmt.Printf("   -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range info.Types {
		fmt.Printf("   - %v\n", t.Type.Name)
	}

	return nil
}
